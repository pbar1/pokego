package battle

var (
	MovePound        = Move{Name: "Pound", Type: TypeNormal, Category: CatPhysical, PP: 35, Power: 40, Accuracy: 100, Effect: NoAdditionalEffect}
	MoveKarateChop   = Move{Name: "Karate Chop", Type: TypeNormal, Category: CatPhysical, PP: 25, Power: 50, Accuracy: 100, Effect: NoAdditionalEffect}
	MoveDoubleSlap   = Move{Name: "Double Slap", Type: TypeNormal, Category: CatPhysical, PP: 10, Power: 15, Accuracy: 85}
	MoveCometPunch   = Move{Name: "Comet Punch", Type: TypeNormal, Category: CatPhysical, PP: 15, Power: 18, Accuracy: 85}
	MoveMegaPunch    = Move{Name: "Mega Punch", Type: TypeNormal, Category: CatPhysical, PP: 20, Power: 80, Accuracy: 85, Effect: NoAdditionalEffect}
	MovePayDay       = Move{Name: "Pay Day", Type: TypeNormal, Category: CatPhysical, PP: 20, Power: 40, Accuracy: 100, Effect: PayDayEffect}
	MoveFirePunch    = Move{Name: "Fire Punch", Type: TypeFire, Category: CatPhysical, PP: 15, Power: 75, Accuracy: 100, Effect: BurnSideEffect1}
	MoveIcePunch     = Move{Name: "Ice Punch", Type: TypeIce, Category: CatPhysical, PP: 15, Power: 75, Accuracy: 100, Effect: FreezeSideEffect}
	MoveThunderPunch = Move{Name: "Thunder Punch", Type: TypeElectric, Category: CatPhysical, PP: 15, Power: 75, Accuracy: 100, Effect: ParalyzeSideEffect1}
	MoveScratch      = Move{Name: "Scratch", Type: TypeNormal, Category: CatPhysical, PP: 35, Power: 40, Accuracy: 100, Effect: NoAdditionalEffect}
	MoveViseGrip     = Move{Name: "Vise Grip", Type: TypeNormal, Category: CatPhysical, PP: 30, Power: 55, Accuracy: 100, Effect: NoAdditionalEffect}
	MoveGuillotine   = Move{Name: "Guillotine", Type: TypeNormal, Category: CatPhysical, PP: 5, Power: -1, Accuracy: 30}
	MoveRazorWind    = Move{Name: "Razor Wind", Type: TypeNormal, Category: CatSpecial, PP: 10, Power: 80, Accuracy: 75}
	MoveSwordsDance  = Move{Name: "Swords Dance", Type: TypeNormal, Category: CatStatus, PP: 30, Power: -1, Accuracy: -1, Effect: AttackUp2Effect}
	MoveCut          = Move{Name: "Cut", Type: TypeNormal, Category: CatPhysical, PP: 30, Power: 50, Accuracy: 95, Effect: NoAdditionalEffect}
	MoveGust         = Move{Name: "Gust", Type: TypeNormal, Category: CatSpecial, PP: 35, Power: 40, Accuracy: 100, Effect: NoAdditionalEffect}
	MoveWingAttack   = Move{Name: "Wing Attack", Type: TypeFlying, Category: CatPhysical, PP: 35, Power: 35, Accuracy: 100, Effect: NoAdditionalEffect}
	MoveWhirlwind    = Move{Name: "Whirlwind", Type: TypeNormal, Category: CatStatus, PP: 20, Power: -1, Accuracy: 85, Effect: SwitchAndTeleportEffect}
	MoveFly          = Move{Name: "Fly", Type: TypeFlying, Category: CatPhysical, PP: 15, Power: 70, Accuracy: 95}
	MoveBind         = Move{Name: "Bind", Type: TypeNormal, Category: CatPhysical, PP: 20, Power: 15, Accuracy: 75}
	MoveSlam         = Move{Name: "Slam", Type: TypeNormal, Category: CatPhysical, PP: 20, Power: 80, Accuracy: 75, Effect: NoAdditionalEffect}
	MoveVineWhip     = Move{Name: "Vine Whip", Type: TypeGrass, Category: CatPhysical, PP: 10, Power: 35, Accuracy: 100, Effect: NoAdditionalEffect}
	MoveStomp        = Move{Name: "Stomp", Type: TypeNormal, Category: CatPhysical, PP: 20, Power: 65, Accuracy: 100}
	MoveDoubleKick   = Move{Name: "Double Kick", Type: TypeFighting, Category: CatPhysical, PP: 30, Power: 30, Accuracy: 100}
	MoveMegaKick     = Move{Name: "Mega Kick", Type: TypeNormal, Category: CatPhysical, PP: 5, Power: 120, Accuracy: 75, Effect: NoAdditionalEffect}
	MoveJumpKick     = Move{Name: "Jump Kick", Type: TypeFighting, Category: CatPhysical, PP: 25, Power: 70, Accuracy: 95}
	MoveRollingKick  = Move{Name: "Rolling Kick", Type: TypeFighting, Category: CatPhysical, PP: 15, Power: 60, Accuracy: 85}
	MoveSandAttack   = Move{Name: "Sand Attack", Type: TypeNormal, Category: CatStatus, PP: 15, Power: -1, Accuracy: 100}
	MoveHeadbutt     = Move{Name: "Headbutt", Type: TypeNormal, Category: CatPhysical, PP: 15, Power: 70, Accuracy: 100}
	MoveHornAttack   = Move{Name: "Horn Attack", Type: TypeNormal, Category: CatPhysical, PP: 25, Power: 65, Accuracy: 100, Effect: NoAdditionalEffect}
	MoveFuryv        = Move{Name: "Fury Attack", Type: TypeNormal, Category: CatPhysical, PP: 20, Power: 15, Accuracy: 85}
	MoveHornDrill    = Move{Name: "Horn Drill", Type: TypeNormal, Category: CatPhysical, PP: 5, Power: -1, Accuracy: 30}
	MoveTackle       = Move{Name: "Tackle", Type: TypeNormal, Category: CatPhysical, PP: 35, Power: 35, Accuracy: 95, Effect: NoAdditionalEffect}
	MoveBodySlam     = Move{Name: "Body Slam", Type: TypeNormal, Category: CatPhysical, PP: 15, Power: 85, Accuracy: 100}
	MoveWrap         = Move{Name: "Wrap", Type: TypeNormal, Category: CatPhysical, PP: 20, Power: 15, Accuracy: 85}
	MoveTakeDown     = Move{Name: "Take Down", Type: TypeNormal, Category: CatPhysical, PP: 20, Power: 90, Accuracy: 85}
	MoveThrash       = Move{Name: "Thrash", Type: TypeNormal, Category: CatPhysical, PP: 20, Power: 90, Accuracy: 100}
	MoveDoubleEdge   = Move{Name: "Double-Edge", Type: TypeNormal, Category: CatPhysical, PP: 15, Power: 100, Accuracy: 100}
	MoveTailWhip     = Move{Name: "Tail Whip", Type: TypeNormal, Category: CatStatus, PP: 30, Power: -1, Accuracy: 100}
	MovePoisonSting  = Move{Name: "Poison Sting", Type: TypePoison, Category: CatPhysical, PP: 35, Power: 15, Accuracy: 100}
	MoveTwineedle    = Move{Name: "Twineedle", Type: TypeBug, Category: CatPhysical, PP: 20, Power: 25, Accuracy: 100}
	MovePinMissile   = Move{Name: "Pin Missile", Type: TypeBug, Category: CatPhysical, PP: 20, Power: 14, Accuracy: 85}
	MoveLeer         = Move{Name: "Leer", Type: TypeNormal, Category: CatStatus, PP: 30, Power: -1, Accuracy: 100}
	MoveBite         = Move{Name: "Bite", Type: TypeNormal, Category: CatPhysical, PP: 25, Power: 60, Accuracy: 100}
	MoveGrowl        = Move{Name: "Growl", Type: TypeNormal, Category: CatStatus, PP: 40, Power: -1, Accuracy: 100}
	MoveRoar         = Move{Name: "Roar", Type: TypeNormal, Category: CatStatus, PP: 20, Power: -1, Accuracy: 100, Effect: SwitchAndTeleportEffect}
	MoveSing         = Move{Name: "Sing", Type: TypeNormal, Category: CatStatus, PP: 15, Power: -1, Accuracy: 55}
	MoveSupersonic   = Move{Name: "Supersonic", Type: TypeNormal, Category: CatStatus, PP: 20, Power: -1, Accuracy: 55}
	MoveSonicBoom    = Move{Name: "Sonic Boom", Type: TypeNormal, Category: CatSpecial, PP: 20, Power: -1, Accuracy: 90}
	MoveDisable      = Move{Name: "Disable", Type: TypeNormal, Category: CatStatus, PP: 20, Power: -1, Accuracy: 55}
	MoveAcid         = Move{Name: "Acid", Type: TypePoison, Category: CatSpecial, PP: 30, Power: 40, Accuracy: 100}
	MoveEmber        = Move{Name: "Ember", Type: TypeFire, Category: CatSpecial, PP: 25, Power: 40, Accuracy: 100}
	MoveFlamethrower = Move{Name: "Flamethrower", Type: TypeFire, Category: CatSpecial, PP: 15, Power: 95, Accuracy: 100}
	MoveMist         = Move{Name: "Mist", Type: TypeIce, Category: CatStatus, PP: 30, Power: -1, Accuracy: -1}
	MoveWaterGun     = Move{Name: "Water Gun", Type: TypeWater, Category: CatSpecial, PP: 25, Power: 40, Accuracy: 100, Effect: NoAdditionalEffect}
	MoveHydroPump    = Move{Name: "Hydro Pump", Type: TypeWater, Category: CatSpecial, PP: 5, Power: 120, Accuracy: 80, Effect: NoAdditionalEffect}
	MoveSurf         = Move{Name: "Surf", Type: TypeWater, Category: CatSpecial, PP: 15, Power: 95, Accuracy: 100, Effect: NoAdditionalEffect}
	MoveIceBeam      = Move{Name: "Ice Beam", Type: TypeIce, Category: CatSpecial, PP: 10, Power: 95, Accuracy: 100}
	MoveBlizzard     = Move{Name: "Blizzard", Type: TypeIce, Category: CatSpecial, PP: 5, Power: 120, Accuracy: 90}
	MovePsybeam      = Move{Name: "Psybeam", Type: TypePsychic, Category: CatSpecial, PP: 20, Power: 65, Accuracy: 100}
	MoveBubbleBeam   = Move{Name: "Bubble Beam", Type: TypeWater, Category: CatSpecial, PP: 20, Power: 65, Accuracy: 100}
	MoveAuroraBeam   = Move{Name: "Aurora Beam", Type: TypeIce, Category: CatSpecial, PP: 20, Power: 65, Accuracy: 100}
	MoveHyperBeam    = Move{Name: "Hyper Beam", Type: TypeNormal, Category: CatSpecial, PP: 5, Power: 150, Accuracy: 90}
	MovePeck         = Move{Name: "Peck", Type: TypeFlying, Category: CatPhysical, PP: 35, Power: 35, Accuracy: 100, Effect: NoAdditionalEffect}
	MoveDrillPeck    = Move{Name: "Drill Peck", Type: TypeFlying, Category: CatPhysical, PP: 20, Power: 80, Accuracy: 100, Effect: NoAdditionalEffect} // TODO: high crit?
	MoveSubmission   = Move{Name: "Submission", Type: TypeFighting, Category: CatPhysical, PP: 25, Power: 80, Accuracy: 80}
	MoveLowKick      = Move{Name: "Low Kick", Type: TypeFighting, Category: CatPhysical, PP: 20, Power: 50, Accuracy: 90}
	MoveCounter      = Move{Name: "Counter", Type: TypeFighting, Category: CatPhysical, PP: 20, Power: -1, Accuracy: 100, Effect: NoAdditionalEffect} // TODO: really?
	MoveSeismicToss  = Move{Name: "Seismic Toss", Type: TypeFighting, Category: CatPhysical, PP: 20, Power: -1, Accuracy: 100}
	MoveStrength     = Move{Name: "Strength", Type: TypeNormal, Category: CatPhysical, PP: 15, Power: 80, Accuracy: 100, Effect: NoAdditionalEffect}
	MoveAbsorb       = Move{Name: "Absorb", Type: TypeGrass, Category: CatSpecial, PP: 20, Power: 20, Accuracy: 100}
	MoveMegaDrain    = Move{Name: "Mega Drain", Type: TypeGrass, Category: CatSpecial, PP: 10, Power: 40, Accuracy: 100}
	MoveLeechSeed    = Move{Name: "Leech Seed", Type: TypeGrass, Category: CatStatus, PP: 10, Power: -1, Accuracy: 90}
	MoveGrowth       = Move{Name: "Growth", Type: TypeNormal, Category: CatStatus, PP: 40, Power: -1, Accuracy: -1}
	MoveRazorLeaf    = Move{Name: "Razor Leaf", Type: TypeGrass, Category: CatPhysical, PP: 25, Power: 55, Accuracy: 95, Effect: NoAdditionalEffect}
	MoveSolarBeam    = Move{Name: "Solar Beam", Type: TypeGrass, Category: CatSpecial, PP: 10, Power: 120, Accuracy: 100}
	MovePoisonPowder = Move{Name: "Poison Powder", Type: TypePoison, Category: CatStatus, PP: 35, Power: -1, Accuracy: 75}
	MoveStunSpore    = Move{Name: "Stun Spore", Type: TypeGrass, Category: CatStatus, PP: 30, Power: -1, Accuracy: 75}
	MoveSleepPowder  = Move{Name: "Sleep Powder", Type: TypeGrass, Category: CatStatus, PP: 15, Power: -1, Accuracy: 75}
	MovePetalDance   = Move{Name: "Petal Dance", Type: TypeGrass, Category: CatSpecial, PP: 20, Power: 70, Accuracy: 100}
	MoveStringShot   = Move{Name: "String Shot", Type: TypeBug, Category: CatStatus, PP: 40, Power: -1, Accuracy: 95}
	MoveDragonRage   = Move{Name: "Dragon Rage", Type: TypeDragon, Category: CatSpecial, PP: 10, Power: -1, Accuracy: 100}
	MoveFireSpin     = Move{Name: "Fire Spin", Type: TypeFire, Category: CatSpecial, PP: 15, Power: 15, Accuracy: 70}
	MoveThunderShock = Move{Name: "Thunder Shock", Type: TypeElectric, Category: CatSpecial, PP: 30, Power: 40, Accuracy: 100}
	MoveThunderbolt  = Move{Name: "Thunderbolt", Type: TypeElectric, Category: CatSpecial, PP: 15, Power: 95, Accuracy: 100}
	MoveThunderWave  = Move{Name: "Thunder Wave", Type: TypeElectric, Category: CatStatus, PP: 20, Power: -1, Accuracy: 100}
	MoveThunder      = Move{Name: "Thunder", Type: TypeElectric, Category: CatSpecial, PP: 10, Power: 120, Accuracy: 70}
	MoveRockThrow    = Move{Name: "Rock Throw", Type: TypeRock, Category: CatPhysical, PP: 15, Power: 50, Accuracy: 65, Effect: NoAdditionalEffect}
	MoveEarthquake   = Move{Name: "Earthquake", Type: TypeGround, Category: CatPhysical, PP: 10, Power: 100, Accuracy: 100, Effect: NoAdditionalEffect}
	MoveFissure      = Move{Name: "Fissure", Type: TypeGround, Category: CatPhysical, PP: 5, Power: -1, Accuracy: 30}
	MoveDig          = Move{Name: "Dig", Type: TypeGround, Category: CatPhysical, PP: 10, Power: 100, Accuracy: 100}
	MoveToxic        = Move{Name: "Toxic", Type: TypePoison, Category: CatStatus, PP: 10, Power: -1, Accuracy: 85}
	MoveConfusion    = Move{Name: "Confusion", Type: TypePsychic, Category: CatSpecial, PP: 25, Power: 50, Accuracy: 100}
	MovePsychic      = Move{Name: "Psychic", Type: TypePsychic, Category: CatSpecial, PP: 10, Power: 90, Accuracy: 100}
	MoveHypnosis     = Move{Name: "Hypnosis", Type: TypePsychic, Category: CatStatus, PP: 20, Power: -1, Accuracy: 60}
	MoveMeditate     = Move{Name: "Meditate", Type: TypePsychic, Category: CatStatus, PP: 40, Power: -1, Accuracy: -1}
	MoveAgility      = Move{Name: "Agility", Type: TypePsychic, Category: CatStatus, PP: 30, Power: -1, Accuracy: -1, Effect: SpeedUp2Effect}
	MoveQuickAttack  = Move{Name: "Quick Attack", Type: TypeNormal, Category: CatPhysical, PP: 30, Power: 40, Accuracy: 100, Effect: NoAdditionalEffect} // TODO: prio
	MoveRage         = Move{Name: "Rage", Type: TypeNormal, Category: CatPhysical, PP: 20, Power: 20, Accuracy: 100}
	MoveTeleport     = Move{Name: "Teleport", Type: TypePsychic, Category: CatStatus, PP: 20, Power: -1, Accuracy: -1}
	MoveNightShade   = Move{Name: "Night Shade", Type: TypeGhost, Category: CatSpecial, PP: 15, Power: -1, Accuracy: 100}
	MoveMimic        = Move{Name: "Mimic", Type: TypeNormal, Category: CatStatus, PP: 10, Power: -1, Accuracy: 100}
	MoveScreech      = Move{Name: "Screech", Type: TypeNormal, Category: CatStatus, PP: 40, Power: -1, Accuracy: 85}
	MoveDoubleTeam   = Move{Name: "Double Team", Type: TypeNormal, Category: CatStatus, PP: 15, Power: -1, Accuracy: -1}
	MoveRecover      = Move{Name: "Recover", Type: TypeNormal, Category: CatStatus, PP: 20, Power: -1, Accuracy: -1}
	MoveHarden       = Move{Name: "Harden", Type: TypeNormal, Category: CatStatus, PP: 30, Power: -1, Accuracy: -1}
	MoveMinimize     = Move{Name: "Minimize", Type: TypeNormal, Category: CatStatus, PP: 20, Power: -1, Accuracy: -1}
	MoveSmokescreen  = Move{Name: "Smokescreen", Type: TypeNormal, Category: CatStatus, PP: 20, Power: -1, Accuracy: 100}
	MoveConfuseRay   = Move{Name: "Confuse Ray", Type: TypeGhost, Category: CatStatus, PP: 10, Power: -1, Accuracy: 100}
	MoveWithdraw     = Move{Name: "Withdraw", Type: TypeWater, Category: CatStatus, PP: 40, Power: -1, Accuracy: -1}
	MoveBaseDefCurl  = Move{Name: "BaseDef Curl", Type: TypeNormal, Category: CatStatus, PP: 40, Power: -1, Accuracy: -1}
	MoveBarrier      = Move{Name: "Barrier", Type: TypePsychic, Category: CatStatus, PP: 30, Power: -1, Accuracy: -1}
	MoveLightScreen  = Move{Name: "Light Screen", Type: TypePsychic, Category: CatStatus, PP: 30, Power: -1, Accuracy: -1}
	MoveHaze         = Move{Name: "Haze", Type: TypeIce, Category: CatStatus, PP: 30, Power: -1, Accuracy: -1}
	MoveReflect      = Move{Name: "Reflect", Type: TypePsychic, Category: CatStatus, PP: 20, Power: -1, Accuracy: -1}
	MoveFocusEnergy  = Move{Name: "Focus Energy", Type: TypeNormal, Category: CatStatus, PP: 30, Power: -1, Accuracy: -1}
	MoveBide         = Move{Name: "Bide", Type: TypeNormal, Category: CatPhysical, PP: 10, Power: -1, Accuracy: 100}
	MoveMetronome    = Move{Name: "Metronome", Type: TypeNormal, Category: CatStatus, PP: 10, Power: -1, Accuracy: -1}
	MoveMirrorMove   = Move{Name: "Mirror Move", Type: TypeFlying, Category: CatStatus, PP: 20, Power: -1, Accuracy: -1}
	MoveSelfDestruct = Move{Name: "Self-Destruct", Type: TypeNormal, Category: CatPhysical, PP: 5, Power: 130, Accuracy: 100}
	MoveEggBomb      = Move{Name: "Egg Bomb", Type: TypeNormal, Category: CatPhysical, PP: 10, Power: 100, Accuracy: 75, Effect: NoAdditionalEffect}
	MoveLick         = Move{Name: "Lick", Type: TypeGhost, Category: CatPhysical, PP: 30, Power: 20, Accuracy: 100}
	MoveSmog         = Move{Name: "Smog", Type: TypePoison, Category: CatSpecial, PP: 20, Power: 20, Accuracy: 70}
	MoveSludge       = Move{Name: "Sludge", Type: TypePoison, Category: CatSpecial, PP: 20, Power: 65, Accuracy: 100}
	MoveBoneClub     = Move{Name: "Bone Club", Type: TypeGround, Category: CatPhysical, PP: 20, Power: 65, Accuracy: 85}
	MoveFireBlast    = Move{Name: "Fire Blast", Type: TypeFire, Category: CatSpecial, PP: 5, Power: 120, Accuracy: 85}
	MoveWaterfall    = Move{Name: "Waterfall", Type: TypeWater, Category: CatPhysical, PP: 15, Power: 80, Accuracy: 100, Effect: NoAdditionalEffect}
	MoveClamp        = Move{Name: "Clamp", Type: TypeWater, Category: CatPhysical, PP: 10, Power: 35, Accuracy: 75}
	MoveSwift        = Move{Name: "Swift", Type: TypeNormal, Category: CatSpecial, PP: 20, Power: 60, Accuracy: -1}
	MoveSkullBash    = Move{Name: "Skull Bash", Type: TypeNormal, Category: CatPhysical, PP: 15, Power: 100, Accuracy: 100}
	MoveSpikeCannon  = Move{Name: "Spike Cannon", Type: TypeNormal, Category: CatPhysical, PP: 15, Power: 20, Accuracy: 100}
	MoveConstrict    = Move{Name: "Constrict", Type: TypeNormal, Category: CatPhysical, PP: 35, Power: 10, Accuracy: 100}
	MoveAmnesia      = Move{Name: "Amnesia", Type: TypePsychic, Category: CatStatus, PP: 20, Power: -1, Accuracy: -1}
	MoveKinesis      = Move{Name: "Kinesis", Type: TypePsychic, Category: CatStatus, PP: 15, Power: -1, Accuracy: 80}
	MoveSoftBoiled   = Move{Name: "Soft-Boiled", Type: TypeNormal, Category: CatStatus, PP: 10, Power: -1, Accuracy: -1}
	MoveHighJumpKick = Move{Name: "High Jump Kick", Type: TypeFighting, Category: CatPhysical, PP: 20, Power: 85, Accuracy: 90}
	MoveGlare        = Move{Name: "Glare", Type: TypeNormal, Category: CatStatus, PP: 30, Power: -1, Accuracy: 75}
	MoveDreamEater   = Move{Name: "Dream Eater", Type: TypePsychic, Category: CatSpecial, PP: 15, Power: 100, Accuracy: 100}
	MovePoisonGas    = Move{Name: "Poison Gas", Type: TypePoison, Category: CatStatus, PP: 40, Power: -1, Accuracy: 55}
	MoveBarrage      = Move{Name: "Barrage", Type: TypeNormal, Category: CatPhysical, PP: 20, Power: 15, Accuracy: 85}
	MoveLeechLife    = Move{Name: "Leech Life", Type: TypeBug, Category: CatPhysical, PP: 15, Power: 20, Accuracy: 100}
	MoveLovelyKiss   = Move{Name: "Lovely Kiss", Type: TypeNormal, Category: CatStatus, PP: 10, Power: -1, Accuracy: 75}
	MoveSkyAttack    = Move{Name: "Sky Attack", Type: TypeFlying, Category: CatPhysical, PP: 5, Power: 140, Accuracy: 90}
	MoveTransform    = Move{Name: "Transform", Type: TypeNormal, Category: CatStatus, PP: 10, Power: -1, Accuracy: -1}
	MoveBubble       = Move{Name: "Bubble", Type: TypeWater, Category: CatSpecial, PP: 30, Power: 20, Accuracy: 100}
	MoveDizzyPunch   = Move{Name: "Dizzy Punch", Type: TypeNormal, Category: CatPhysical, PP: 10, Power: 70, Accuracy: 100, Effect: NoAdditionalEffect}
	MoveSpore        = Move{Name: "Spore", Type: TypeGrass, Category: CatStatus, PP: 15, Power: -1, Accuracy: 100}
	MoveFlash        = Move{Name: "Flash", Type: TypeNormal, Category: CatStatus, PP: 20, Power: -1, Accuracy: 70}
	MovePsywave      = Move{Name: "Psywave", Type: TypePsychic, Category: CatSpecial, PP: 15, Power: -1, Accuracy: 80}
	MoveSplash       = Move{Name: "Splash", Type: TypeNormal, Category: CatStatus, PP: 40, Power: -1, Accuracy: -1}
	MoveAcidArmor    = Move{Name: "Acid Armor", Type: TypePoison, Category: CatStatus, PP: 40, Power: -1, Accuracy: -1}
	MoveCrabhammer   = Move{Name: "Crabhammer", Type: TypeWater, Category: CatPhysical, PP: 10, Power: 90, Accuracy: 85, Effect: NoAdditionalEffect}
	MoveExplosion    = Move{Name: "Explosion", Type: TypeNormal, Category: CatPhysical, PP: 5, Power: 170, Accuracy: 100}
	MoveFurySwipes   = Move{Name: "Fury Swipes", Type: TypeNormal, Category: CatPhysical, PP: 15, Power: 18, Accuracy: 80}
	MoveBonemerang   = Move{Name: "Bonemerang", Type: TypeGround, Category: CatPhysical, PP: 10, Power: 50, Accuracy: 90}
	MoveRest         = Move{Name: "Rest", Type: TypePsychic, Category: CatStatus, PP: 10, Power: -1, Accuracy: -1}
	MoveRockSlide    = Move{Name: "Rock Slide", Type: TypeRock, Category: CatPhysical, PP: 10, Power: 75, Accuracy: 90, Effect: NoAdditionalEffect}
	MoveHyperFang    = Move{Name: "Hyper Fang", Type: TypeNormal, Category: CatPhysical, PP: 15, Power: 80, Accuracy: 90}
	MoveSharpen      = Move{Name: "Sharpen", Type: TypeNormal, Category: CatStatus, PP: 30, Power: -1, Accuracy: -1}
	MoveConversion   = Move{Name: "Conversion", Type: TypeNormal, Category: CatStatus, PP: 30, Power: -1, Accuracy: -1}
	MoveTriAttack    = Move{Name: "Tri Attack", Type: TypeNormal, Category: CatSpecial, PP: 10, Power: 80, Accuracy: 100, Effect: NoAdditionalEffect}
	MoveSuperFang    = Move{Name: "Super Fang", Type: TypeNormal, Category: CatPhysical, PP: 10, Power: -1, Accuracy: 90}
	MoveSlash        = Move{Name: "Slash", Type: TypeNormal, Category: CatPhysical, PP: 20, Power: 70, Accuracy: 100, Effect: NoAdditionalEffect} // TODO: high crit
	MoveSubstitute   = Move{Name: "Substitute", Type: TypeNormal, Category: CatStatus, PP: 10, Power: -1, Accuracy: -1}
	MoveStruggle     = Move{Name: "Struggle", Type: TypeNormal, Category: CatPhysical, PP: 10, Power: 50, Accuracy: 100}
)