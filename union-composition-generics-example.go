package main

type flying interface {
	fly() string
}

type Bird struct {
	wingcolor string
}
func (bird Bird)fly() string {
	return "bird flying with wingcolor " + bird.wingcolor
}

type Squirrel struct {
	tailcolor string
}
func (squirrel Squirrel)fly() string {
	return "squirrel flying with tailcolor " + squirrel.tailcolor
}

type running interface {
	run() string
}

type Lion struct {
	manecolor string
}
func (lion Lion)run() string {
	return "lion running with manecolor " + lion.manecolor
}

type Cheeta struct {
	furcolor string
}
func (cheeta Cheeta)run() string {
	return "cheeta running with furcolor " + cheeta.furcolor
}

type animal interface {
	Bird | Squirrel | Lion | Cheeta
}

func testanimal[T animal](a T) {
	switch any(a).(type) {
	case flying:
		println(any(a).(flying).fly())
	case running:
		println(any(a).(running).run())
	}
}

func main() {
	testanimal(Bird{"red"})
	testanimal(Squirrel{"green"})
	testanimal(Lion{"blue"})
	testanimal(Cheeta{"yellow"})
}
