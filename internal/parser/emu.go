package parser

import (
	"fmt"
)

var systems = map[string]string{
	"3do":             "3DO",
	"ags":             "Adventure Game Studio Game Engine",
	"amiga":           "Commodore Amiga",
	"amiga1200":       "Commodore Amiga 1200",
	"amiga600":        "Commodore Amiga 600",
	"amigacd32":       "Commodore Amiga CD32",
	"amstradcpc":      "Amstrad CPC",
	"android":         "Google Android",
	"apple2":          "Apple II",
	"apple2gs":        "Apple IIGS",
	"arcade":          "Arcade",
	"astrocde":        "Bally Astrocade",
	"atari2600":       "Atari 2600",
	"atari5200":       "Atari 5200",
	"atari7800":       "Atari 7800 ProSystem",
	"atari800":        "Atari 800",
	"atarijaguar":     "Atari Jaguar",
	"atarijaguarcd":   "Atari Jaguar CD",
	"atarilynx":       "Atari Lynx",
	"atarist":         "Atari ST",
	"atarixe":         "Atari XE",
	"atomiswave":      "Atomiswave",
	"bbcmicro":        "BBC Micro",
	"c64":             "Commodore 64",
	"cavestory":       "Cave Story (NXEngine)",
	"cdimono1":        "Philips CD-i",
	"cdtv":            "Commodore CDTV",
	"chailove":        "ChaiLove Game Engine",
	"channelf":        "Fairchild Channel F",
	"coco":            "Tandy Color Computer",
	"colecovision":    "ColecoVision",
	"cps":             "Capcom Play System",
	"daphne":          "Daphne Arcade LaserDisc Emulator",
	"desktop":         "Desktop Applications",
	"doom":            "Doom",
	"dos":             "DOS (PC)",
	"dragon32":        "Dragon 32",
	"dreamcast":       "Sega Dreamcast",
	"easyrpg":         "EasyRPG Game Engine",
	"epic":            "Epic Games Store",
	"famicom":         "Nintendo Family Computer",
	"fba":             "FinalBurn Alpha",
	"fbneo":           "FinalBurn Neo",
	"fds":             "Nintendo Famicom Disk System",
	"flash":           "Adobe Flash",
	"fmtowns":         "Fujitsu FM Towns",
	"gameandwatch":    "Nintendo Game and Watch",
	"gamegear":        "Sega Game Gear",
	"gb":              "Nintendo Game Boy",
	"gba":             "Nintendo Game Boy Advance",
	"gbc":             "Nintendo Game Boy Color",
	"gc":              "Nintendo GameCube",
	"genesis":         "Sega Genesis",
	"gx4000":          "Amstrad GX4000",
	"intellivision":   "Mattel Electronics Intellivision",
	"j2me":            "Java 2 Micro Edition (J2ME)",
	"kodi":            "Kodi Home Theatre Software",
	"lutris":          "Lutris Open Gaming Platform",
	"lutro":           "Lutro Game Engine",
	"macintosh":       "Apple Macintosh",
	"mame":            "Multiple Arcade Machine Emulator",
	"mame-advmame":    "AdvanceMAME",
	"mame-mame4all":   "MAME4ALL",
	"mastersystem":    "Sega Master System",
	"megacd":          "Sega Mega-CD",
	"megacdjp":        "Sega Mega-CD",
	"megadrive":       "Sega Mega Drive",
	"megaduck":        "Creatronic Mega Duck",
	"mess":            "Multi Emulator Super System",
	"model2":          "Sega Model 2",
	"model3":          "Sega Model 3",
	"moonlight":       "Moonlight Game Streaming",
	"moto":            "Thomson MO/TO Series",
	"msx":             "MSX",
	"msx1":            "MSX1",
	"msx2":            "MSX2",
	"msxturbor":       "MSX Turbo R",
	"mugen":           "M.U.G.E.N Game Engine",
	"multivision":     "Othello Multivision",
	"n3ds":            "Nintendo 3DS",
	"n64":             "Nintendo 64",
	"n64dd":           "Nintendo 64DD",
	"naomi":           "Sega NAOMI",
	"naomigd":         "Sega NAOMI GD-ROM",
	"nds":             "Nintendo DS",
	"neogeo":          "SNK Neo Geo",
	"neogeocd":        "SNK Neo Geo CD",
	"neogeocdjp":      "SNK Neo Geo CD",
	"nes":             "Nintendo Entertainment System",
	"ngp":             "SNK Neo Geo Pocket",
	"ngpc":            "SNK Neo Geo Pocket Color",
	"odyssey2":        "Magnavox Odyssey2",
	"openbor":         "OpenBOR Game Engine",
	"oric":            "Tangerine Computer Systems Oric",
	"palm":            "Palm OS",
	"pc":              "IBM PC",
	"pc88":            "NEC PC-8800 Series",
	"pc98":            "NEC PC-9800 Series",
	"pcengine":        "NEC PC Engine",
	"pcenginecd":      "NEC PC Engine CD",
	"pcfx":            "NEC PC-FX",
	"pico8":           "PICO-8 Fantasy Console",
	"pokemini":        "Nintendo Pokémon Mini",
	"ports":           "Ports",
	"ps2":             "Sony PlayStation 2",
	"ps3":             "Sony PlayStation 3",
	"ps4":             "Sony PlayStation 4",
	"psp":             "Sony PlayStation Portable",
	"psvita":          "Sony PlayStation Vita",
	"psx":             "Sony PlayStation",
	"samcoupe":        "SAM Coupé",
	"satellaview":     "Nintendo Satellaview",
	"saturn":          "Sega Saturn",
	"saturnjp":        "Sega Saturn",
	"scummvm":         "ScummVM Game Engine",
	"sega32x":         "Sega Mega Drive 32X",
	"sega32xjp":       "Sega Super 32X",
	"sega32xna":       "Sega Genesis 32X",
	"segacd":          "Sega CD",
	"sfc":             "Nintendo SFC (Super Famicom)",
	"sg-1000":         "Sega SG-1000",
	"sgb":             "Nintendo Super Game Boy",
	"snes":            "Nintendo SNES (Super Nintendo)",
	"snesna":          "Nintendo SNES (Super Nintendo)",
	"solarus":         "Solarus Game Engine",
	"spectravideo":    "Spectravideo",
	"steam":           "Valve Steam",
	"stratagus":       "Stratagus Game Engine",
	"sufami":          "Bandai SuFami Turbo",
	"supergrafx":      "NEC SuperGrafx",
	"supervision":     "Watara Supervision",
	"switch":          "Nintendo Switch",
	"symbian":         "Symbian",
	"tanodragon":      "Tano Dragon",
	"tg-cd":           "NEC TurboGrafx-CD",
	"tg16":            "NEC TurboGrafx-16",
	"ti99":            "Texas Instruments TI-99",
	"tic80":           "TIC-80 Game Engine",
	"to8":             "Thomson TO8",
	"trs-80":          "Tandy TRS-80",
	"uzebox":          "Uzebox",
	"vectrex":         "Vectrex",
	"vic20":           "Commodore VIC-20",
	"videopac":        "Philips Videopac G7000",
	"virtualboy":      "Nintendo Virtual Boy",
	"wii":             "Nintendo Wii",
	"wiiu":            "Nintendo Wii U",
	"wiiu/roms ":      "Nintendo Wii U (custom system)",
	"wonderswan":      "Bandai WonderSwan",
	"wonderswancolor": "Bandai WonderSwan Color",
	"x1":              "Sharp X1",
	"x68000":          "Sharp X68000",
	"xbox":            "Microsoft Xbox",
	"xbox360":         "Microsoft Xbox 360",
	"zmachine":        "Infocom Z-machine",
	"zx81":            "Sinclair ZX81",
	"zxspectrum":      "Sinclair ZX Spectrum",
}

func ValidateSystem(ident string) (string, error) {
	platform, ok := systems[ident]
	if !ok {
		return "", fmt.Errorf("invalid system identifier %s", ident)
	}

	return platform, nil
}
