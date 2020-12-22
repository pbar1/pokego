package battle

type (
	Generation      uint8
	DamageType      uint8
	DamageCategory  uint8
	StatusCondition uint8
	MoveEffect      uint8

	Move struct {
		Name     string
		Type     DamageType
		Category DamageCategory
		PP       int
		Power    int
		Accuracy int
		Effect   func(*Move, *ActivePokemon, *Board) bool
	}

	BasePokemon struct {
		DexNo   int
		Name    string
		Type1   DamageType
		Type2   DamageType
		BaseHP  int
		BaseAtk int
		BaseDef int
		BaseSpe int
		BaseSpc int
	}

	// https://bulbapedia.bulbagarden.net/wiki/Pok%C3%A9mon_data_structure_in_Generation_I
	BattlePokemon struct {
		DexNo        int
		CurrentHP    int
		Level        int
		Status       StatusCondition
		Type1        DamageType
		Type2        DamageType
		Move1        Move
		Move2        Move
		Move3        Move
		Move4        Move
		StatHP       int
		StatAtk      int
		StatDef      int
		StatSpe      int
		StatSpc      int
		SleepCounter int
	}

	NewBattlePokemonOpts struct {
		Base       BasePokemon
		Level      int
		DVAtk      int
		DVDef      int
		DVSpe      int
		DVSpc      int
		StatExpHP  int
		StatExpAtk int
		StatExpDef int
		StatExpSpe int
		StatExpSpc int
		Move1      Move
		Move2      Move
		Move3      Move
		Move4      Move
	}

	ActivePokemon struct {
		Pokemon             BattlePokemon
		LvAtk               int
		LvDef               int
		LvSpe               int
		LvSpc               int
		LvAcc               int
		LvEvasion           int
		ConfusionCounter    int
		ToxSeedCounter      int
		SubstituteCurrentHP int
		ReflectUp           bool
		LightScreenUp       bool
		Flinch              bool
	}

	Player struct {
		Party  [6]BattlePokemon
		Active ActivePokemon
	}

	Board struct {
		Player1 Player
		Player2 Player
	}
)

const (
	Gen1 Generation = iota
	Gen2
	Gen3
	Gen4
	Gen5
	Gen6
	Gen7
	Gen8

	TypeNone DamageType = iota
	TypeNormal
	TypeFighting
	TypeFlying
	TypePoison
	TypeGround
	TypeRock
	TypeBug
	TypeGhost
	TypeSteel
	TypeFire
	TypeWater
	TypeGrass
	TypeElectric
	TypePsychic
	TypeIce
	TypeDragon
	TypeDark
	TypeFairy

	CatNone DamageCategory = iota
	CatPhysical
	CatSpecial
	CatStatus

	StatusNone StatusCondition = iota
	StatusSleep
	StatusPoison
	StatusBurn
	StatusFreeze
	StatusParalyze
	StatusFaint
)
