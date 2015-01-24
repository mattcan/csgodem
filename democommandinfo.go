package main

const (
	FDEMO_NORMAL            int32 = 0
	FDEMO_USE_ORIGIN2       int32 = (1 << 0)
	FDEMO_USE_ANGLES2       int32 = (1 << 1)
	FDEMO_NOINTERP          int32 = (1 << 2)
	MAX_SPLITSCREEN_CLIENTS int32 = 2
)

// QAngle setup
type QAngle struct {
	x, y, z float32
}

func (q *QAngle) InitDefault() {
	q.x = 0.0
	q.y = 0.0
	q.z = 0.0
}

func (q *QAngle) Init(x float32, y float32, z float32) {
	q.x = x
	q.y = y
	q.z = z
}

// Vector setup
type Vector struct {
	x, y, z float32
}

func (v *Vector) InitDefault() {
	q.x = 0.0
	q.y = 0.0
	q.z = 0.0
}

func (v *Vector) Init(x float32, y float32, z float32) {
	q.x = x
	q.y = y
	q.z = z
}

// Split type
type Split struct {
	flags int32

	viewOrigin      Vector
	viewAngles      QAngle
	localViewAngles QAngle

	viewOrigin2      Vector
	viewAngles2      QAngle
	localViewAngles2 QAngle
}

// Set default Split values
func (s *Split) Init() {
	s.flags = FDEMO_NORMAL

	s.viewOrigin.Init()
	s.viewAngles.Init()
	s.localViewAngles.Init()

	s.viewOrigin2.Init()
	s.viewAngles2.Init()
	s.localViewAngles.Init()
}

// Copy the src Split to the current one
func (s *Split) Copy(src Split) {
	s.flags = src.flags

	s.viewOrigin = src.viewOrigin
	s.viewAngles = src.viewAngles
	s.localViewAngles = src.localViewAngles

	s.viewOrigin2 = src.viewOrigin2
	s.viewAngles2 = src.viewAngles2
	s.localViewAngles2 = src.localViewAngles2
}

func (s *Split) GetViewOrigin() *Vector {
	if flags & FDEMO_USE_ORIGIN2 {
		return s.viewOrigin2
	}
	return s.viewOrigin
}

func (s *Split) GetViewAngles() *QAngle {
	if flags & FDEMO_USE_ANGLES2 {
		return s.viewAngles2
	}
	return s.viewAngles
}

func (s *Split) GetLocalViewAngles() *QAngle {
	if flags & FDEMO_USE_ANGLES2 {
		return s.localViewAngles2
	}
	return s.localViewAngles
}

func (s *Split) Reset() {
	s.flags = 0
	s.viewOrigin2 = s.viewOrigin
	s.viewAngles2 = s.viewAngles
	s.localViewAngles2 = s.localViewAngles
}

type DemoCmdInfo struct {
	u [MAX_SPLITSCREEN_CLIENTS]Split
}

func (d *DemoCmdInfo) Reset() {
	for i := 0; i < MAX_SPLITSCREEN_CLIENTS; i++ {
		d.u[i].Reset()
	}
}
