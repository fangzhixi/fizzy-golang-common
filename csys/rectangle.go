package csys

import (
	"errors"
	"fmt"
	"math"
)

//坐标
type Coord struct {
	X int32 `protobuf:"zigzag32,1,opt,name=x,proto3" json:"x,omitempty"` //横坐标
	Y int32 `protobuf:"zigzag32,2,opt,name=y,proto3" json:"y,omitempty"` //纵坐标
}

//文本的坐标，以四个顶点坐标表示 注意：此字段可能返回 null，表示取不到有效值
type Polygon struct {
	LeftTop     *Coord `protobuf:"bytes,1,opt,name=left_top,json=leftTop,proto3" json:"left_top,omitempty"`             // 左上顶点坐标
	RightTop    *Coord `protobuf:"bytes,2,opt,name=right_top,json=rightTop,proto3" json:"right_top,omitempty"`          // 右上顶点坐标
	RightBottom *Coord `protobuf:"bytes,3,opt,name=right_bottom,json=rightBottom,proto3" json:"right_bottom,omitempty"` // 右下顶点坐标
	LeftBottom  *Coord `protobuf:"bytes,4,opt,name=left_bottom,json=leftBottom,proto3" json:"left_bottom,omitempty"`    // 左下顶点坐标
}

//矩形坐标
type Rect struct {
	X      int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`           //左上角x
	Y      int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`           //左上角y
	Width  int32 `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`   //宽度
	Height int32 `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"` //高度
}

// 矩形坐标转换
func RectToPolygon(rect *Rect, angle float64) (*Polygon, error) {
	/*
		STEP_1: 将Rect转成Polygon
		STEP_2: 将Polygon与角度结合为新Polygon
	*/
	if rect == nil {
		return nil, errors.New("坐标值为空")
	}
	//STEP_1
	var (
		polygon = &Polygon{
			LeftTop:     &Coord{X: rect.X, Y: rect.Y},
			LeftBottom:  &Coord{X: rect.X, Y: rect.Y + rect.Height},
			RightTop:    &Coord{X: rect.X + rect.Width, Y: rect.Y},
			RightBottom: &Coord{X: rect.X + rect.Width, Y: rect.Y + rect.Height},
		}
		originCoord = &Coord{
			X: rect.X + rect.Width/2,
			Y: rect.Y + rect.Height/2,
		} //旋转原点
	)
	//STEP_2
	polygon, err := RectangleRotation(polygon, originCoord, angle)
	return polygon, err
}

// 二维系四边形旋转
func RectangleRotation(rectangle *Polygon, originCoord *Coord, angle float64) (polygon *Polygon, err error) {
	/*
		STEP_1: 将四边形坐标Polygon 拆解成四角坐标
		STEP_2: 依次计算四点坐标旋转后的坐标值
	*/
	if rectangle == nil ||
		rectangle.LeftTop == nil ||
		rectangle.RightTop == nil ||
		rectangle.LeftBottom == nil ||
		rectangle.RightBottom == nil ||
		originCoord == nil {
		return nil, errors.New("矩形值不能为空")
	}
	var errStr = "二维系坐标旋转失败"
	polygon = &Polygon{}
	polygon.LeftTop, err = CoordRotation(rectangle.LeftTop, originCoord, angle)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", errStr, err)
	}
	polygon.RightTop, err = CoordRotation(rectangle.RightTop, originCoord, angle)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", errStr, err)
	}
	polygon.RightBottom, err = CoordRotation(rectangle.RightBottom, originCoord, angle)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", errStr, err)
	}
	polygon.LeftBottom, err = CoordRotation(rectangle.LeftBottom, originCoord, angle)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", errStr, err)
	}
	return polygon, nil
}

/*
二维系坐标旋转
parameter:
	targetCoord 目标点
	originCoord 旋转原点
	angle		旋转角度
return:
	coord 旋转后的坐标
	err
*/
func CoordRotation(targetCoord, originCoord *Coord, angle float64) (*Coord, error) {
	/*
		算法步骤:
			STEP_1: 计算目标点targetCoord与旋转原点originCoord连成直线与X轴的弧度R_1
				注: R = arctan y/x
			STEP_2: 以旋转原点为中心建立笛卡尔坐标系, 对目标点targetCoord分类讨论
			STEP_3: 将传入角度值angle换算成弧度R_2
				注: 角度∠A换算弧度R -> ∠A = R * 180 / π
			STEP_4: 将R_1与R_2相加得到最终旋转的弧度
			STEP_5: 计算目标点与旋转原点的距离, 作为旋转半径D
			STEP_6: 通过余弦定理计算旋转后的坐标值
	*/
	if targetCoord == nil || originCoord == nil {
		return nil, errors.New("坐标值为空")
	}
	var (
		xRand float64 //目标点与X轴弧度
		rand  float64 //旋转弧度
		r     float64 //目标点与原点距离
	)
	//STEP_1
	if float64(originCoord.X-targetCoord.X) == 0.00 {
		xRand = 90 * math.Pi / 180 //tan90 -> +∞
	} else {
		xRand = math.Atan(float64(originCoord.Y-targetCoord.Y) / float64(originCoord.X-targetCoord.X))
	}
	//STEP_2
	if originCoord.Y-targetCoord.Y > 0 && originCoord.X-targetCoord.X < 0 {
		xRand = math.Pi + xRand //第一象限
	} else if originCoord.Y-targetCoord.Y < 0 && originCoord.X-targetCoord.X > 0 {
		xRand = 2*math.Pi + xRand //第三象限
	} else if originCoord.Y-targetCoord.Y < 0 && originCoord.X-targetCoord.X < 0 {
		xRand = math.Pi + xRand //第四象限
	}
	//STEP_3
	rand = angle * math.Pi / 180 //角度换算弧度
	//STEP_4
	rand = xRand + rand //从X轴顺时针旋转弧度 = 目标点与X轴弧度 + 旋转弧度
	//STEP_5
	r = math.Sqrt(math.Pow(float64(targetCoord.X-originCoord.X), 2) + math.Pow(float64(targetCoord.Y-originCoord.Y), 2))
	//STEP_6
	targetCoord = &Coord{
		X: originCoord.X - int32(math.Cos(rand)*r),
		Y: originCoord.Y - int32(math.Sin(rand)*r),
	}
	return targetCoord, nil
}

// 四边形包含
func IsContainPolygon(bigPolygon, smallPolygon *Polygon) (err error) {
	/*
		四边形包含原理: 小四边形四角坐标点均在大四边形内则存在包含关系
			STEP_1: 将四边形坐标Polygon 拆解成四角坐标
			STEP_2: 依次判断四角坐标是否在大四边形内
	*/
	var errStr = "四边形不包含"
	if bigPolygon == nil ||
		smallPolygon == nil {
		return errors.New("矩形值不能为空")
	}
	if err = IsCoordInsidePolygon(bigPolygon, smallPolygon.LeftTop); err != nil {
		return fmt.Errorf("%s: %v", errStr, err)
	}
	if err = IsCoordInsidePolygon(bigPolygon, smallPolygon.LeftBottom); err != nil {
		return fmt.Errorf("%s: %v", errStr, err)
	}
	if err = IsCoordInsidePolygon(bigPolygon, smallPolygon.RightTop); err != nil {
		return fmt.Errorf("%s: %v", errStr, err)
	}
	if err = IsCoordInsidePolygon(bigPolygon, smallPolygon.RightBottom); err != nil {
		return fmt.Errorf("%s: %v", errStr, err)
	}
	return nil
}

// 判断坐标点是否在四边形内
func IsCoordInsidePolygon(polygon *Polygon, coord *Coord) (err error) {
	/*
		原理: 对于不规则图形，可以通过射线法判断，即计算射线与多边形各边的交点，如果是偶数，则点在多边形外，否则在多边形内
			STEP_1: 对四边形拆解成四条边
			STEP_2: 统计射线与四边交点数
			STEP_3: 判断统计值是否为奇数, 是则返回正确, 否则返回异常
	*/
	if polygon == nil ||
		polygon.LeftTop == nil ||
		polygon.RightTop == nil ||
		polygon.LeftBottom == nil ||
		polygon.RightBottom == nil ||
		coord == nil {
		return errors.New("坐标值为空")
	}
	//STEP_1
	//STEP_2
	var count = 0 //记录从点发出的射线与边相交数
	//记录射线与线段交点
	if err = isRayWithLineIntersection(polygon.LeftTop, polygon.RightTop, coord); err == nil {
		count++
	}
	if err = isRayWithLineIntersection(polygon.RightTop, polygon.RightBottom, coord); err == nil {
		count++
	}
	if err = isRayWithLineIntersection(polygon.RightBottom, polygon.LeftBottom, coord); err == nil {
		count++
	}
	if err = isRayWithLineIntersection(polygon.LeftBottom, polygon.LeftTop, coord); err == nil {
		count++
	}
	//STEP_3
	if count%2 == 1 {
		return nil
	} else {
		return fmt.Errorf("坐标点不在四边形内: %v", err)
	}
}

// 判断从点发出射线是否与线段相交 (默认坐标点平行X轴向右发出射线)
func isRayWithLineIntersection(endpointA, endpointB, coord *Coord) (err error) {
	/*
		算法原理:
			1.从t坐标点沿X轴方向发出射线
			2.t点与线段(v1-v2)要发生相交，t.y必须在线段的两个顶点的y值之间，即  t.y<v2.y&&t.y>v1.y
			3.满足上面这个条件以后，只需要判断该点在线段的左侧还是右侧，如果在线段左侧，则该射线与线段相交，
			要判断t在左侧还是右侧，需要先求得水平线与线段的交点c的x坐标：c.x=(t.y-v1.y)*(v2.x-v1.x)/(v2.y-v1.y)+v1.x
			由上两条，可以推得，t点与线段相交的条件为： t.y<v2.y  &&  t.y>v1.y  &&  c.x<((t.y-v1.y)*(v2.x-v1.x)/(v2.y-v1.y)+v1.x)
		实现步骤:
			STEP_1: 判断目标点Y值是否在线段Y值区间
			STEP_2: 求线段一元函数y = k*x + b
			STEP_3: 求射线与线段交点c
			STEP_4: 判断交点c是否在目标点右侧, 是则返回正确, 否则返回异常

	*/
	if endpointA == nil || endpointB == nil || coord == nil {
		return errors.New("坐标点为空")
	}
	var (
		//y=k*x+b
		k             float64
		b             float64
		intersectionX int32
	)
	//STEP_1
	if (coord.Y >= endpointA.Y && coord.Y >= endpointB.Y) || //大于等于保证两线段交点只与射线相交一次
		(endpointA.Y > coord.Y && endpointB.Y > coord.Y) {
		return errors.New("点线不相交")
	}
	//STEP_2
	if endpointB.X-endpointA.X == 0 || float64(endpointB.Y-endpointA.Y)/float64(endpointB.X-endpointA.X) == 0 {
		//STEP_3
		intersectionX = endpointB.X //k -> +∞ => x = y - b
	} else {
		k = float64(endpointB.Y-endpointA.Y) / float64(endpointB.X-endpointA.X)
		b = float64(endpointA.Y) - k*float64(endpointA.X)
		//STEP_3
		intersectionX = int32((float64(coord.Y) - b) / k)
	}
	//STEP_4
	if intersectionX > coord.X {
		return nil
	}
	return errors.New("点线不相交")
}
