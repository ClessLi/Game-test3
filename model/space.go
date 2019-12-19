package model

import (
	"fmt"
	"github.com/ClessLi/Game-test/resource"
	"github.com/ClessLi/Game-test/sprite"
	"github.com/ClessLi/resolvForGame/resolv"
)

type Space []Shape

// NewSpace creates a new Space for shapes to exist in and be tested against in.
func NewSpace() *Space {
	sp := &Space{}
	return sp
}

// Add adds the designated Shapes to the Space. You cannot add the Space to itself.
func (sp *Space) Add(shapes ...Shape) {
	for _, shape := range shapes {
		if shape == sp {
			panic(fmt.Sprintf("ERROR! Space %s cannot add itself!", shape))
		}
		*sp = append(*sp, shape)
	}
}

// Remove removes the designated Shapes from the Space.
func (sp *Space) Remove(shapes ...Shape) {

	for _, shape := range shapes {

		for deleteIndex, s := range *sp {

			if s == shape {
				s := *sp
				s[deleteIndex] = nil
				s = append(s[:deleteIndex], s[deleteIndex+1:]...)
				*sp = s
				break
			}

		}

	}

}

// Clear "resets" the Space, cleaning out the Space of references to Shapes.
func (sp *Space) Clear() {
	*sp = make(Space, 0)
}

// IsColliding returns whether the provided Shape is colliding with something in this Space.
func (sp *Space) IsColliding(shape Shape) bool {

	for _, other := range *sp {

		if other != shape {

			if shape.IsColliding(other) {
				return true
			}

		}

	}

	return false

}

// GetCollidingShapes returns a Space comprised of Shapes that collide with the checking Shape.
func (sp *Space) GetCollidingShapes(shape Shape) *Space {

	newSpace := NewSpace()

	for _, other := range *sp {
		if other != shape {
			if shape.IsColliding(other) {
				newSpace.Add(other)
			}
		}
	}

	return newSpace

}

// Resolve runs Resolve() using the checking Shape, checking against all other Shapes in the Space. The first Collision
// that returns true is the Collision that gets returned.
func (sp *Space) Resolve(checkingShape Shape, deltaX, deltaY int32) Collision {

	res := Collision{}

	for _, other := range *sp {

		if other != checkingShape && checkingShape.WouldBeColliding(other, deltaX, deltaY) {
			res = Resolve(checkingShape, other, deltaX, deltaY)
			if res.Colliding() {
				break
			}
		}

	}

	return res

}

// Filter filters out a Space, returning a new Space comprised of Shapes that return true for the boolean function you provide.
// This can be used to focus on a set of object for collision testing or resolution, or lower the number of Shapes to test
// by filtering some out beforehand.
func (sp *Space) Filter(filterFunc func(Shape) bool) *Space {
	subSpace := NewSpace()
	for _, shape := range *sp {
		if filterFunc(shape) {
			subSpace.Add(shape)
		}
	}
	return subSpace
}

// FilterByTags filters a Space out, creating a new Space that has just the Shapes that have all of the specified tags.
func (sp *Space) FilterByTags(tags ...string) *Space {
	return sp.Filter(func(s Shape) bool {
		if s.HasTags(tags...) {
			return true
		}
		return false
	})
}

// FilterOutByTags filters a Space out, creating a new Space that has just the Shapes that don't have all of the specified tags.
func (sp *Space) FilterOutByTags(tags ...string) *Space {
	return sp.Filter(func(s Shape) bool {
		if s.HasTags(tags...) {
			return false
		}
		return true
	})
}

// Contains returns true if the Shape provided exists within the Space.
func (sp *Space) Contains(shape Shape) bool {
	for _, s := range *sp {
		if s == shape {
			return true
		}
	}
	return false
}

func (sp *Space) String() string {
	str := ""
	for _, s := range *sp {
		str += fmt.Sprintf("%v   ", s)
	}
	return str
}

/* -----------------------------
   --  SPACE-SHAPE FUNCTIONS  --
   -----------------------------
These functions allows a Space to fulfill the contract of a Shape as well, thereby allowing them to serve as easy-use
compound Shapes themselves. Functions that should logically function on all Shapes within a Space do that, while functions
that return singular values look at the first shape as a "root" of sorts.
*/

// WouldBeColliding returns true if any of the Shapes within the Space would be colliding should they move along the delta
// X and Y values provided (dx and dy).
func (sp *Space) WouldBeColliding(other Shape, dx, dy int32) bool {

	for _, shape := range *sp {

		if shape == other {
			return false
		}

		if shape.WouldBeColliding(other, dx, dy) {
			return true
		}

	}

	return false

}

// GetTags returns the tag list of the first Shape within the Space. If there are no Shapes within the Space,
// it returns an empty array of string type.
func (sp *Space) GetTags() []string {
	if len(*sp) > 0 {
		return (*sp)[0].GetTags()
	}
	return []string{}
}

// AddTags sets the provided tags on all Shapes contained within the Space.
func (sp *Space) AddTags(tags ...string) {
	for _, shape := range *sp {
		shape.AddTags(tags...)
	}
}

// RemoveTags removes the provided tags from all Shapes contained within the Space.
func (sp *Space) RemoveTags(tags ...string) {
	for _, shape := range *sp {
		shape.RemoveTags(tags...)
	}
}

// ClearTags removes all tags from all Shapes within the Space.
func (sp *Space) ClearTags() {
	for _, shape := range *sp {
		shape.ClearTags()
	}
}

// HasTags returns true if all of the Shapes contained within the Space have the tags specified.
func (sp *Space) HasTags(tags ...string) bool {

	for _, shape := range *sp {
		if !shape.HasTags(tags...) {
			return false
		}
	}
	return true

}

// GetData returns the pointer to the object contained in the Data field of the first Shape within the Space. If there aren't
// any Shapes within the Space, it returns nil.
func (sp *Space) GetData() interface{} {

	if len(*sp) > 0 {
		return (*sp)[0].GetData()
	}
	return nil

}

// SetData sets the pointer provided to the Data field of all Shapes within the Space.
func (sp *Space) SetData(data interface{}) {

	for _, shape := range *sp {
		shape.SetData(data)
	}

}

// GetXY returns the X and Y position of the first Shape in the Space. If there aren't any Shapes within the Space, it
// returns 0, 0.
func (sp *Space) GetXY() (int32, int32) {

	if len(*sp) > 0 {
		return (*sp)[0].GetXY()
	}
	return 0, 0

}

func (sp *Space) GetXY2() (int32, int32) {

	if len(*sp) > 0 {
		return (*sp)[0].GetXY2()
	}
	return 0, 0

}

// SetXY sets the X and Y position of all Shapes within the Space to the position provided using the first Shape's position as
// reference. Basically, it moves the first Shape within the Space to the target location and then moves all other Shapes
// by the same delta movement.
func (sp *Space) SetXY(x, y int32) {

	if len(*sp) > 0 {

		x0, y0 := sp.GetXY()
		dx := x - x0
		dy := y - y0

		for _, shape := range *sp {
			shape.Move(dx, dy)
		}

	}

}

// Move moves all Shapes in the Space by the displacement provided.
func (sp *Space) Move(dx, dy int32) {
	for _, shape := range *sp {
		shape.Move(dx, dy)
	}
}

// Length returns the length of the Space (number of Shapes contained within the Space). This is a convenience function, standing in for len(*space).
func (sp *Space) Length() int {
	return len(*sp)
}

// Get allows you to get a Shape by index from the Space easily. This is a convenience function, standing in for (*space)[index].
func (sp *Space) Get(index int) Shape {
	return (*sp)[index]
}

func (sp *Space) Draw(renderer *sprite.SpriteRenderer) {
	// TO-DO: 图像集合空间渲染方法，目前暂在gameMap中Draw方法根据camera来判定集合空间内图像对象是否被渲染。
}

func (sp *Space) GetShapeObj() resolv.Shape {
	if len(*sp) > 0 {
		return (*sp)[0].GetShapeObj()
	}
	return nil
}

func (sp *Space) SetTexture(t *resource.Texture2D) {

}

func (sp *Space) GetFriction() float32 {
	if len(*sp) > 0 {
		return (*sp)[0].GetFriction()
	}
	return 0.0
}

func (sp *Space) SetFriction(friction float32) {
	if len(*sp) > 0 {

		for _, shape := range *sp {
			shape.SetFriction(friction)
		}

	}
}

func (sp *Space) GetMaxSpd() float32 {
	if len(*sp) > 0 {
		return (*sp)[0].GetMaxSpd()
	}
	return 0.0
}

func (sp *Space) SetMaxSpd(spd float32) {
	if len(*sp) > 0 {

		for _, shape := range *sp {
			shape.SetMaxSpd(spd)
		}

	}
}
