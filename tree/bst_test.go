package tree

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBSTChecker(t *testing.T) {
	Convey("Valid tree with left and right nodes", t, func() {
		root := &Node{
			value: 10,
		}
		left := &Node{
			value: 5,
		}
		right := &Node{
			value: 15,
		}
		root.left = left
		root.right = right

		//So(root.ToString(), ShouldEqual, "5 10 15")
		So(root.Max(), ShouldEqual, 15)
		So(root.SecondMax(), ShouldEqual, 10)
		//So(IsBST(root), ShouldBeTrue)
	})
	Convey("Left Childis bigger than root", t, func() {
		root := &Node{
			value: 10,
		}
		left := &Node{
			value: 11,
		}
		right := &Node{
			value: 15,
		}
		root.left = left
		root.right = right

		So(IsBST(root), ShouldBeFalse)
	})
	Convey("Valid tree with two levels", t, func() {
		/*
		      10
		    /     |
		  5        20
		/  |       /
	   2	7    15  25
		 */
		root := &Node{
			value: 10,
		}
		left := &Node{
			value: 5,
		}
		right := &Node{
			value: 15,
		}
		root.left = left
		root.right = right

		//level 2
		left1 := &Node{
			value: 2,
		}
		right1 := &Node{
			value: 7,
		}
		left.left = left1
		left.right = right1

		left2 := &Node{
			value: 15,
		}
		right2 := &Node{
			value: 25,
		}
		right.left = left2
		right.right = right2

		So(IsBST(root), ShouldBeTrue)
		So(root.Max(), ShouldEqual, 25)
	})
	Convey("Failure - Left 11 is big than left", t, func() {
		/*
		      10
		    /     |
		  5        20
		/  |       /
	   6	7    15  25
		 */
		root := &Node{
			value: 10,
		}
		left := &Node{
			value: 5,
		}
		right := &Node{
			value: 15,
		}
		root.left = left
		root.right = right

		//level 2
		left1 := &Node{
			value: 6,
		}
		right1 := &Node{
			value: 7,
		}
		left.left = left1
		left.right = right1

		left2 := &Node{
			value: 15,
		}
		right2 := &Node{
			value: 25,
		}
		right.left = left2
		right.right = right2

		So(IsBST(root), ShouldBeFalse)
	})
	Convey("Failure - right 22 is less than right", t, func() {
		/*
		      10
		    /     |
		  5        20
		/  |       /
	   2	7    15  18
		 */
		root := &Node{
			value: 10,
		}
		left := &Node{
			value: 2,
		}
		right := &Node{
			value: 15,
		}
		root.left = left
		root.right = right

		//level 2
		left1 := &Node{
			value: 6,
		}
		right1 := &Node{
			value: 7,
		}
		left.left = left1
		left.right = right1

		left2 := &Node{
			value: 15,
		}
		right2 := &Node{
			value: 25,
		}
		right.left = left2
		right.right = right2

		So(IsBST(root), ShouldBeFalse)
	})
}

func TestGetSecondMax(t *testing.T) {
	/*
			  ( 5 )
			/     \
		  (3)     (8)
		 /  \     /  \
	   (1)  (4) (7)  (9)
	 */
	Convey("node has right subtree", t, func() {
		root := &Node{
			value: 5,
		}
		left := &Node{
			value: 3,
		}
		right := &Node{
			value: 8,
		}
		root.left = left
		root.right = right

		//level 2
		left1 := &Node{
			value: 1,
		}
		right1 := &Node{
			value: 4,
		}
		left.left = left1
		left.right = right1

		left2 := &Node{
			value: 7,
		}
		right2 := &Node{
			value: 9,
		}
		right.left = left2
		right.right = right2

		So(root.SecondMax(), ShouldEqual, 8)
	})
	/*
          ( 5 )
        /     \
      (3)     (8)
     /  \     /  \
   (1)  (4) (7)  (12)
                 /
               (10)
               /  \
             (9)  (11)
	*/
	Convey("node doesn't have right subtree -> find max in left subtree", t, func() {
		root := &Node{
			value: 5,
		}
		left := &Node{
			value: 3,
		}
		right := &Node{
			value: 8,
		}
		root.left = left
		root.right = right

		//level 2
		left1 := &Node{
			value: 1,
		}
		right1 := &Node{
			value: 4,
		}
		left.left = left1
		left.right = right1

		left2 := &Node{
			value: 7,
		}
		right2 := &Node{
			value: 12,
		}
		right.left = left2
		right.right = right2

		//level 3
		left31 := &Node{
			value: 10,
		}
		right2.left = left31

		//level 4
		left41 := &Node{
			value: 9,
		}
		right41 := &Node{
			value: 11,
		}
		left31.left = left41
		left31.right = right41

		So(root.SecondMax(), ShouldEqual, 11)
	})
}