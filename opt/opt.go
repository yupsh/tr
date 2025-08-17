package opt

// Boolean flag types with constants
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

// Flags represents the configuration options for the tr command
type Flags struct {
	Delete     DeleteFlag     // Delete characters in set1
	Squeeze    SqueezeFlag    // Squeeze multiple consecutive characters
	Complement ComplementFlag // Use complement of set1
}

// Configure methods for the opt system
func (f DeleteFlag) Configure(flags *Flags) { flags.Delete = f }
func (f SqueezeFlag) Configure(flags *Flags) { flags.Squeeze = f }
func (f ComplementFlag) Configure(flags *Flags) { flags.Complement = f }
