package command

type DeleteFlag bool

const (
	Delete   DeleteFlag = true
	NoDelete DeleteFlag = false
)

type SqueezeFlag bool

const (
	Squeeze   SqueezeFlag = true
	NoSqueeze SqueezeFlag = false
)

type ComplementFlag bool

const (
	Complement   ComplementFlag = true
	NoComplement ComplementFlag = false
)

type flags struct {
	Delete     DeleteFlag
	Squeeze    SqueezeFlag
	Complement ComplementFlag
}

func (f DeleteFlag) Configure(flags *flags)     { flags.Delete = f }
func (f SqueezeFlag) Configure(flags *flags)    { flags.Squeeze = f }
func (f ComplementFlag) Configure(flags *flags) { flags.Complement = f }
