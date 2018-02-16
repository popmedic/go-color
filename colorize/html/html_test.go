package html

import (
	"testing"

	"github.com/popmedic/go-color/colorize"
)

func TestFgAliceBlue(t *testing.T) { test("color:AliceBlue;", FgAliceBlue(), t) }
func TestBgAliceBlue(t *testing.T) { test("background-color:AliceBlue;", BgAliceBlue(), t) }

func TestFgAntiqueWhite(t *testing.T) { test("color:AntiqueWhite;", FgAntiqueWhite(), t) }
func TestBgAntiqueWhite(t *testing.T) { test("background-color:AntiqueWhite;", BgAntiqueWhite(), t) }

func TestFgAqua(t *testing.T) { test("color:Aqua;", FgAqua(), t) }
func TestBgAqua(t *testing.T) { test("background-color:Aqua;", BgAqua(), t) }

func TestFgAquamarine(t *testing.T) { test("color:Aquamarine;", FgAquamarine(), t) }
func TestBgAquamarine(t *testing.T) { test("background-color:Aquamarine;", BgAquamarine(), t) }

func TestFgAzure(t *testing.T) { test("color:Azure;", FgAzure(), t) }
func TestBgAzure(t *testing.T) { test("background-color:Azure;", BgAzure(), t) }

func TestFgBeige(t *testing.T) { test("color:Beige;", FgBeige(), t) }
func TestBgBeige(t *testing.T) { test("background-color:Beige;", BgBeige(), t) }

func TestFgBisque(t *testing.T) { test("color:Bisque;", FgBisque(), t) }
func TestBgBisque(t *testing.T) { test("background-color:Bisque;", BgBisque(), t) }

func TestFgBlack(t *testing.T) { test("color:Black;", FgBlack(), t) }
func TestBgBlack(t *testing.T) { test("background-color:Black;", BgBlack(), t) }

func TestFgBlanchedAlmond(t *testing.T) { test("color:BlanchedAlmond;", FgBlanchedAlmond(), t) }
func TestBgBlanchedAlmond(t *testing.T) {
	test("background-color:BlanchedAlmond;", BgBlanchedAlmond(), t)
}

func TestFgBlue(t *testing.T) { test("color:Blue;", FgBlue(), t) }
func TestBgBlue(t *testing.T) { test("background-color:Blue;", BgBlue(), t) }

func TestFgBlueViolet(t *testing.T) { test("color:BlueViolet;", FgBlueViolet(), t) }
func TestBgBlueViolet(t *testing.T) { test("background-color:BlueViolet;", BgBlueViolet(), t) }

func TestFgBrown(t *testing.T) { test("color:Brown;", FgBrown(), t) }
func TestBgBrown(t *testing.T) { test("background-color:Brown;", BgBrown(), t) }

func TestFgBurlyWood(t *testing.T) { test("color:BurlyWood;", FgBurlyWood(), t) }
func TestBgBurlyWood(t *testing.T) { test("background-color:BurlyWood;", BgBurlyWood(), t) }

func TestFgCadetBlue(t *testing.T) { test("color:CadetBlue;", FgCadetBlue(), t) }
func TestBgCadetBlue(t *testing.T) { test("background-color:CadetBlue;", BgCadetBlue(), t) }

func TestFgChartreuse(t *testing.T) { test("color:Chartreuse;", FgChartreuse(), t) }
func TestBgChartreuse(t *testing.T) { test("background-color:Chartreuse;", BgChartreuse(), t) }

func TestFgChocolate(t *testing.T) { test("color:Chocolate;", FgChocolate(), t) }
func TestBgChocolate(t *testing.T) { test("background-color:Chocolate;", BgChocolate(), t) }

func TestFgCoral(t *testing.T) { test("color:Coral;", FgCoral(), t) }
func TestBgCoral(t *testing.T) { test("background-color:Coral;", BgCoral(), t) }

func TestFgCornflowerBlue(t *testing.T) { test("color:CornflowerBlue;", FgCornflowerBlue(), t) }
func TestBgCornflowerBlue(t *testing.T) {
	test("background-color:CornflowerBlue;", BgCornflowerBlue(), t)
}

func TestFgCornsilk(t *testing.T) { test("color:Cornsilk;", FgCornsilk(), t) }
func TestBgCornsilk(t *testing.T) { test("background-color:Cornsilk;", BgCornsilk(), t) }

func TestFgCrimson(t *testing.T) { test("color:Crimson;", FgCrimson(), t) }
func TestBgCrimson(t *testing.T) { test("background-color:Crimson;", BgCrimson(), t) }

func TestFgCyan(t *testing.T) { test("color:Cyan;", FgCyan(), t) }
func TestBgCyan(t *testing.T) { test("background-color:Cyan;", BgCyan(), t) }

func TestFgDarkBlue(t *testing.T) { test("color:DarkBlue;", FgDarkBlue(), t) }
func TestBgDarkBlue(t *testing.T) { test("background-color:DarkBlue;", BgDarkBlue(), t) }

func TestFgDarkCyan(t *testing.T) { test("color:DarkCyan;", FgDarkCyan(), t) }
func TestBgDarkCyan(t *testing.T) { test("background-color:DarkCyan;", BgDarkCyan(), t) }

func TestFgDarkGoldenRod(t *testing.T) { test("color:DarkGoldenRod;", FgDarkGoldenRod(), t) }
func TestBgDarkGoldenRod(t *testing.T) { test("background-color:DarkGoldenRod;", BgDarkGoldenRod(), t) }

func TestFgDarkGray(t *testing.T) { test("color:DarkGray;", FgDarkGray(), t) }
func TestBgDarkGray(t *testing.T) { test("background-color:DarkGray;", BgDarkGray(), t) }

func TestFgDarkGrey(t *testing.T) { test("color:DarkGrey;", FgDarkGrey(), t) }
func TestBgDarkGrey(t *testing.T) { test("background-color:DarkGrey;", BgDarkGrey(), t) }

func TestFgDarkGreen(t *testing.T) { test("color:DarkGreen;", FgDarkGreen(), t) }
func TestBgDarkGreen(t *testing.T) { test("background-color:DarkGreen;", BgDarkGreen(), t) }

func TestFgDarkKhaki(t *testing.T) { test("color:DarkKhaki;", FgDarkKhaki(), t) }
func TestBgDarkKhaki(t *testing.T) { test("background-color:DarkKhaki;", BgDarkKhaki(), t) }

func TestFgDarkMagenta(t *testing.T) { test("color:DarkMagenta;", FgDarkMagenta(), t) }
func TestBgDarkMagenta(t *testing.T) { test("background-color:DarkMagenta;", BgDarkMagenta(), t) }

func TestFgDarkOliveGreen(t *testing.T) { test("color:DarkOliveGreen;", FgDarkOliveGreen(), t) }
func TestBgDarkOliveGreen(t *testing.T) {
	test("background-color:DarkOliveGreen;", BgDarkOliveGreen(), t)
}

func TestFgDarkOrange(t *testing.T) { test("color:DarkOrange;", FgDarkOrange(), t) }
func TestBgDarkOrange(t *testing.T) { test("background-color:DarkOrange;", BgDarkOrange(), t) }

func TestFgDarkOrchid(t *testing.T) { test("color:DarkOrchid;", FgDarkOrchid(), t) }
func TestBgDarkOrchid(t *testing.T) { test("background-color:DarkOrchid;", BgDarkOrchid(), t) }

func TestFgDarkRed(t *testing.T) { test("color:DarkRed;", FgDarkRed(), t) }
func TestBgDarkRed(t *testing.T) { test("background-color:DarkRed;", BgDarkRed(), t) }

func TestFgDarkSalmon(t *testing.T) { test("color:DarkSalmon;", FgDarkSalmon(), t) }
func TestBgDarkSalmon(t *testing.T) { test("background-color:DarkSalmon;", BgDarkSalmon(), t) }

func TestFgDarkSeaGreen(t *testing.T) { test("color:DarkSeaGreen;", FgDarkSeaGreen(), t) }
func TestBgDarkSeaGreen(t *testing.T) { test("background-color:DarkSeaGreen;", BgDarkSeaGreen(), t) }

func TestFgDarkSlateBlue(t *testing.T) { test("color:DarkSlateBlue;", FgDarkSlateBlue(), t) }
func TestBgDarkSlateBlue(t *testing.T) { test("background-color:DarkSlateBlue;", BgDarkSlateBlue(), t) }

func TestFgDarkSlateGray(t *testing.T) { test("color:DarkSlateGray;", FgDarkSlateGray(), t) }
func TestBgDarkSlateGray(t *testing.T) { test("background-color:DarkSlateGray;", BgDarkSlateGray(), t) }

func TestFgDarkSlateGrey(t *testing.T) { test("color:DarkSlateGrey;", FgDarkSlateGrey(), t) }
func TestBgDarkSlateGrey(t *testing.T) { test("background-color:DarkSlateGrey;", BgDarkSlateGrey(), t) }

func TestFgDarkTurquoise(t *testing.T) { test("color:DarkTurquoise;", FgDarkTurquoise(), t) }
func TestBgDarkTurquoise(t *testing.T) { test("background-color:DarkTurquoise;", BgDarkTurquoise(), t) }

func TestFgDarkViolet(t *testing.T) { test("color:DarkViolet;", FgDarkViolet(), t) }
func TestBgDarkViolet(t *testing.T) { test("background-color:DarkViolet;", BgDarkViolet(), t) }

func TestFgDeepPink(t *testing.T) { test("color:DeepPink;", FgDeepPink(), t) }
func TestBgDeepPink(t *testing.T) { test("background-color:DeepPink;", BgDeepPink(), t) }

func TestFgDeepSkyBlue(t *testing.T) { test("color:DeepSkyBlue;", FgDeepSkyBlue(), t) }
func TestBgDeepSkyBlue(t *testing.T) { test("background-color:DeepSkyBlue;", BgDeepSkyBlue(), t) }

func TestFgDimGray(t *testing.T) { test("color:DimGray;", FgDimGray(), t) }
func TestBgDimGray(t *testing.T) { test("background-color:DimGray;", BgDimGray(), t) }

func TestFgDimGrey(t *testing.T) { test("color:DimGrey;", FgDimGrey(), t) }
func TestBgDimGrey(t *testing.T) { test("background-color:DimGrey;", BgDimGrey(), t) }

func TestFgDodgerBlue(t *testing.T) { test("color:DodgerBlue;", FgDodgerBlue(), t) }
func TestBgDodgerBlue(t *testing.T) { test("background-color:DodgerBlue;", BgDodgerBlue(), t) }

func TestFgFireBrick(t *testing.T) { test("color:FireBrick;", FgFireBrick(), t) }
func TestBgFireBrick(t *testing.T) { test("background-color:FireBrick;", BgFireBrick(), t) }

func TestFgFloralWhite(t *testing.T) { test("color:FloralWhite;", FgFloralWhite(), t) }
func TestBgFloralWhite(t *testing.T) { test("background-color:FloralWhite;", BgFloralWhite(), t) }

func TestFgForestGreen(t *testing.T) { test("color:ForestGreen;", FgForestGreen(), t) }
func TestBgForestGreen(t *testing.T) { test("background-color:ForestGreen;", BgForestGreen(), t) }

func TestFgFuchsia(t *testing.T) { test("color:Fuchsia;", FgFuchsia(), t) }
func TestBgFuchsia(t *testing.T) { test("background-color:Fuchsia;", BgFuchsia(), t) }

func TestFgGainsboro(t *testing.T) { test("color:Gainsboro;", FgGainsboro(), t) }
func TestBgGainsboro(t *testing.T) { test("background-color:Gainsboro;", BgGainsboro(), t) }

func TestFgGhostWhite(t *testing.T) { test("color:GhostWhite;", FgGhostWhite(), t) }
func TestBgGhostWhite(t *testing.T) { test("background-color:GhostWhite;", BgGhostWhite(), t) }

func TestFgGold(t *testing.T) { test("color:Gold;", FgGold(), t) }
func TestBgGold(t *testing.T) { test("background-color:Gold;", BgGold(), t) }

func TestFgGoldenRod(t *testing.T) { test("color:GoldenRod;", FgGoldenRod(), t) }
func TestBgGoldenRod(t *testing.T) { test("background-color:GoldenRod;", BgGoldenRod(), t) }

func TestFgGray(t *testing.T) { test("color:Gray;", FgGray(), t) }
func TestBgGray(t *testing.T) { test("background-color:Gray;", BgGray(), t) }

func TestFgGrey(t *testing.T) { test("color:Grey;", FgGrey(), t) }
func TestBgGrey(t *testing.T) { test("background-color:Grey;", BgGrey(), t) }

func TestFgGreen(t *testing.T) { test("color:Green;", FgGreen(), t) }
func TestBgGreen(t *testing.T) { test("background-color:Green;", BgGreen(), t) }

func TestFgGreenYellow(t *testing.T) { test("color:GreenYellow;", FgGreenYellow(), t) }
func TestBgGreenYellow(t *testing.T) { test("background-color:GreenYellow;", BgGreenYellow(), t) }

func TestFgHoneyDew(t *testing.T) { test("color:HoneyDew;", FgHoneyDew(), t) }
func TestBgHoneyDew(t *testing.T) { test("background-color:HoneyDew;", BgHoneyDew(), t) }

func TestFgHotPink(t *testing.T) { test("color:HotPink;", FgHotPink(), t) }
func TestBgHotPink(t *testing.T) { test("background-color:HotPink;", BgHotPink(), t) }

func TestFgIndianRed(t *testing.T) { test("color:IndianRed ;", FgIndianRed(), t) }
func TestBgIndianRed(t *testing.T) { test("background-color:IndianRed ;", BgIndianRed(), t) }

func TestFgIndigo(t *testing.T) { test("color:Indigo ;", FgIndigo(), t) }
func TestBgIndigo(t *testing.T) { test("background-color:Indigo ;", BgIndigo(), t) }

func TestFgIvory(t *testing.T) { test("color:Ivory;", FgIvory(), t) }
func TestBgIvory(t *testing.T) { test("background-color:Ivory;", BgIvory(), t) }

func TestFgKhaki(t *testing.T) { test("color:Khaki;", FgKhaki(), t) }
func TestBgKhaki(t *testing.T) { test("background-color:Khaki;", BgKhaki(), t) }

func TestFgLavender(t *testing.T) { test("color:Lavender;", FgLavender(), t) }
func TestBgLavender(t *testing.T) { test("background-color:Lavender;", BgLavender(), t) }

func TestFgLavenderBlush(t *testing.T) { test("color:LavenderBlush;", FgLavenderBlush(), t) }
func TestBgLavenderBlush(t *testing.T) { test("background-color:LavenderBlush;", BgLavenderBlush(), t) }

func TestFgLawnGreen(t *testing.T) { test("color:LawnGreen;", FgLawnGreen(), t) }
func TestBgLawnGreen(t *testing.T) { test("background-color:LawnGreen;", BgLawnGreen(), t) }

func TestFgLemonChiffon(t *testing.T) { test("color:LemonChiffon;", FgLemonChiffon(), t) }
func TestBgLemonChiffon(t *testing.T) { test("background-color:LemonChiffon;", BgLemonChiffon(), t) }

func TestFgLightBlue(t *testing.T) { test("color:LightBlue;", FgLightBlue(), t) }
func TestBgLightBlue(t *testing.T) { test("background-color:LightBlue;", BgLightBlue(), t) }

func TestFgLightCoral(t *testing.T) { test("color:LightCoral;", FgLightCoral(), t) }
func TestBgLightCoral(t *testing.T) { test("background-color:LightCoral;", BgLightCoral(), t) }

func TestFgLightCyan(t *testing.T) { test("color:LightCyan;", FgLightCyan(), t) }
func TestBgLightCyan(t *testing.T) { test("background-color:LightCyan;", BgLightCyan(), t) }

func TestFgLightGoldenRodYellow(t *testing.T) {
	test("color:LightGoldenRodYellow;", FgLightGoldenRodYellow(), t)
}
func TestBgLightGoldenRodYellow(t *testing.T) {
	test("background-color:LightGoldenRodYellow;", BgLightGoldenRodYellow(), t)
}

func TestFgLightGray(t *testing.T) { test("color:LightGray;", FgLightGray(), t) }
func TestBgLightGray(t *testing.T) { test("background-color:LightGray;", BgLightGray(), t) }

func TestFgLightGrey(t *testing.T) { test("color:LightGrey;", FgLightGrey(), t) }
func TestBgLightGrey(t *testing.T) { test("background-color:LightGrey;", BgLightGrey(), t) }

func TestFgLightGreen(t *testing.T) { test("color:LightGreen;", FgLightGreen(), t) }
func TestBgLightGreen(t *testing.T) { test("background-color:LightGreen;", BgLightGreen(), t) }

func TestFgLightPink(t *testing.T) { test("color:LightPink;", FgLightPink(), t) }
func TestBgLightPink(t *testing.T) { test("background-color:LightPink;", BgLightPink(), t) }

func TestFgLightSalmon(t *testing.T) { test("color:LightSalmon;", FgLightSalmon(), t) }
func TestBgLightSalmon(t *testing.T) { test("background-color:LightSalmon;", BgLightSalmon(), t) }

func TestFgLightSeaGreen(t *testing.T) { test("color:LightSeaGreen;", FgLightSeaGreen(), t) }
func TestBgLightSeaGreen(t *testing.T) { test("background-color:LightSeaGreen;", BgLightSeaGreen(), t) }

func TestFgLightSkyBlue(t *testing.T) { test("color:LightSkyBlue;", FgLightSkyBlue(), t) }
func TestBgLightSkyBlue(t *testing.T) { test("background-color:LightSkyBlue;", BgLightSkyBlue(), t) }

func TestFgLightSlateGray(t *testing.T) { test("color:LightSlateGray;", FgLightSlateGray(), t) }
func TestBgLightSlateGray(t *testing.T) {
	test("background-color:LightSlateGray;", BgLightSlateGray(), t)
}

func TestFgLightSlateGrey(t *testing.T) { test("color:LightSlateGrey;", FgLightSlateGrey(), t) }
func TestBgLightSlateGrey(t *testing.T) {
	test("background-color:LightSlateGrey;", BgLightSlateGrey(), t)
}

func TestFgLightSteelBlue(t *testing.T) { test("color:LightSteelBlue;", FgLightSteelBlue(), t) }
func TestBgLightSteelBlue(t *testing.T) {
	test("background-color:LightSteelBlue;", BgLightSteelBlue(), t)
}

func TestFgLightYellow(t *testing.T) { test("color:LightYellow;", FgLightYellow(), t) }
func TestBgLightYellow(t *testing.T) { test("background-color:LightYellow;", BgLightYellow(), t) }

func TestFgLime(t *testing.T) { test("color:Lime;", FgLime(), t) }
func TestBgLime(t *testing.T) { test("background-color:Lime;", BgLime(), t) }

func TestFgLimeGreen(t *testing.T) { test("color:LimeGreen;", FgLimeGreen(), t) }
func TestBgLimeGreen(t *testing.T) { test("background-color:LimeGreen;", BgLimeGreen(), t) }

func TestFgLinen(t *testing.T) { test("color:Linen;", FgLinen(), t) }
func TestBgLinen(t *testing.T) { test("background-color:Linen;", BgLinen(), t) }

func TestFgMagenta(t *testing.T) { test("color:Magenta;", FgMagenta(), t) }
func TestBgMagenta(t *testing.T) { test("background-color:Magenta;", BgMagenta(), t) }

func TestFgMaroon(t *testing.T) { test("color:Maroon;", FgMaroon(), t) }
func TestBgMaroon(t *testing.T) { test("background-color:Maroon;", BgMaroon(), t) }

func TestFgMediumAquaMarine(t *testing.T) { test("color:MediumAquaMarine;", FgMediumAquaMarine(), t) }
func TestBgMediumAquaMarine(t *testing.T) {
	test("background-color:MediumAquaMarine;", BgMediumAquaMarine(), t)
}

func TestFgMediumBlue(t *testing.T) { test("color:MediumBlue;", FgMediumBlue(), t) }
func TestBgMediumBlue(t *testing.T) { test("background-color:MediumBlue;", BgMediumBlue(), t) }

func TestFgMediumOrchid(t *testing.T) { test("color:MediumOrchid;", FgMediumOrchid(), t) }
func TestBgMediumOrchid(t *testing.T) { test("background-color:MediumOrchid;", BgMediumOrchid(), t) }

func TestFgMediumPurple(t *testing.T) { test("color:MediumPurple;", FgMediumPurple(), t) }
func TestBgMediumPurple(t *testing.T) { test("background-color:MediumPurple;", BgMediumPurple(), t) }

func TestFgMediumSeaGreen(t *testing.T) { test("color:MediumSeaGreen;", FgMediumSeaGreen(), t) }
func TestBgMediumSeaGreen(t *testing.T) {
	test("background-color:MediumSeaGreen;", BgMediumSeaGreen(), t)
}

func TestFgMediumSlateBlue(t *testing.T) { test("color:MediumSlateBlue;", FgMediumSlateBlue(), t) }
func TestBgMediumSlateBlue(t *testing.T) {
	test("background-color:MediumSlateBlue;", BgMediumSlateBlue(), t)
}

func TestFgMediumSpringGreen(t *testing.T) { test("color:MediumSpringGreen;", FgMediumSpringGreen(), t) }
func TestBgMediumSpringGreen(t *testing.T) {
	test("background-color:MediumSpringGreen;", BgMediumSpringGreen(), t)
}

func TestFgMediumTurquoise(t *testing.T) { test("color:MediumTurquoise;", FgMediumTurquoise(), t) }
func TestBgMediumTurquoise(t *testing.T) {
	test("background-color:MediumTurquoise;", BgMediumTurquoise(), t)
}

func TestFgMediumVioletRed(t *testing.T) { test("color:MediumVioletRed;", FgMediumVioletRed(), t) }
func TestBgMediumVioletRed(t *testing.T) {
	test("background-color:MediumVioletRed;", BgMediumVioletRed(), t)
}

func TestFgMidnightBlue(t *testing.T) { test("color:MidnightBlue;", FgMidnightBlue(), t) }
func TestBgMidnightBlue(t *testing.T) { test("background-color:MidnightBlue;", BgMidnightBlue(), t) }

func TestFgMintCream(t *testing.T) { test("color:MintCream;", FgMintCream(), t) }
func TestBgMintCream(t *testing.T) { test("background-color:MintCream;", BgMintCream(), t) }

func TestFgMistyRose(t *testing.T) { test("color:MistyRose;", FgMistyRose(), t) }
func TestBgMistyRose(t *testing.T) { test("background-color:MistyRose;", BgMistyRose(), t) }

func TestFgMoccasin(t *testing.T) { test("color:Moccasin;", FgMoccasin(), t) }
func TestBgMoccasin(t *testing.T) { test("background-color:Moccasin;", BgMoccasin(), t) }

func TestFgNavajoWhite(t *testing.T) { test("color:NavajoWhite;", FgNavajoWhite(), t) }
func TestBgNavajoWhite(t *testing.T) { test("background-color:NavajoWhite;", BgNavajoWhite(), t) }

func TestFgNavy(t *testing.T) { test("color:Navy;", FgNavy(), t) }
func TestBgNavy(t *testing.T) { test("background-color:Navy;", BgNavy(), t) }

func TestFgOldLace(t *testing.T) { test("color:OldLace;", FgOldLace(), t) }
func TestBgOldLace(t *testing.T) { test("background-color:OldLace;", BgOldLace(), t) }

func TestFgOlive(t *testing.T) { test("color:Olive;", FgOlive(), t) }
func TestBgOlive(t *testing.T) { test("background-color:Olive;", BgOlive(), t) }

func TestFgOliveDrab(t *testing.T) { test("color:OliveDrab;", FgOliveDrab(), t) }
func TestBgOliveDrab(t *testing.T) { test("background-color:OliveDrab;", BgOliveDrab(), t) }

func TestFgOrange(t *testing.T) { test("color:Orange;", FgOrange(), t) }
func TestBgOrange(t *testing.T) { test("background-color:Orange;", BgOrange(), t) }

func TestFgOrangeRed(t *testing.T) { test("color:OrangeRed;", FgOrangeRed(), t) }
func TestBgOrangeRed(t *testing.T) { test("background-color:OrangeRed;", BgOrangeRed(), t) }

func TestFgOrchid(t *testing.T) { test("color:Orchid;", FgOrchid(), t) }
func TestBgOrchid(t *testing.T) { test("background-color:Orchid;", BgOrchid(), t) }

func TestFgPaleGoldenRod(t *testing.T) { test("color:PaleGoldenRod;", FgPaleGoldenRod(), t) }
func TestBgPaleGoldenRod(t *testing.T) { test("background-color:PaleGoldenRod;", BgPaleGoldenRod(), t) }

func TestFgPaleGreen(t *testing.T) { test("color:PaleGreen;", FgPaleGreen(), t) }
func TestBgPaleGreen(t *testing.T) { test("background-color:PaleGreen;", BgPaleGreen(), t) }

func TestFgPaleTurquoise(t *testing.T) { test("color:PaleTurquoise;", FgPaleTurquoise(), t) }
func TestBgPaleTurquoise(t *testing.T) { test("background-color:PaleTurquoise;", BgPaleTurquoise(), t) }

func TestFgPaleVioletRed(t *testing.T) { test("color:PaleVioletRed;", FgPaleVioletRed(), t) }
func TestBgPaleVioletRed(t *testing.T) { test("background-color:PaleVioletRed;", BgPaleVioletRed(), t) }

func TestFgPapayaWhip(t *testing.T) { test("color:PapayaWhip;", FgPapayaWhip(), t) }
func TestBgPapayaWhip(t *testing.T) { test("background-color:PapayaWhip;", BgPapayaWhip(), t) }

func TestFgPeachPuff(t *testing.T) { test("color:PeachPuff;", FgPeachPuff(), t) }
func TestBgPeachPuff(t *testing.T) { test("background-color:PeachPuff;", BgPeachPuff(), t) }

func TestFgPeru(t *testing.T) { test("color:Peru;", FgPeru(), t) }
func TestBgPeru(t *testing.T) { test("background-color:Peru;", BgPeru(), t) }

func TestFgPink(t *testing.T) { test("color:Pink;", FgPink(), t) }
func TestBgPink(t *testing.T) { test("background-color:Pink;", BgPink(), t) }

func TestFgPlum(t *testing.T) { test("color:Plum;", FgPlum(), t) }
func TestBgPlum(t *testing.T) { test("background-color:Plum;", BgPlum(), t) }

func TestFgPowderBlue(t *testing.T) { test("color:PowderBlue;", FgPowderBlue(), t) }
func TestBgPowderBlue(t *testing.T) { test("background-color:PowderBlue;", BgPowderBlue(), t) }

func TestFgPurple(t *testing.T) { test("color:Purple;", FgPurple(), t) }
func TestBgPurple(t *testing.T) { test("background-color:Purple;", BgPurple(), t) }

func TestFgRebeccaPurple(t *testing.T) { test("color:RebeccaPurple;", FgRebeccaPurple(), t) }
func TestBgRebeccaPurple(t *testing.T) { test("background-color:RebeccaPurple;", BgRebeccaPurple(), t) }

func TestFgRed(t *testing.T) { test("color:Red;", FgRed(), t) }
func TestBgRed(t *testing.T) { test("background-color:Red;", BgRed(), t) }

func TestFgRosyBrown(t *testing.T) { test("color:RosyBrown;", FgRosyBrown(), t) }
func TestBgRosyBrown(t *testing.T) { test("background-color:RosyBrown;", BgRosyBrown(), t) }

func TestFgRoyalBlue(t *testing.T) { test("color:RoyalBlue;", FgRoyalBlue(), t) }
func TestBgRoyalBlue(t *testing.T) { test("background-color:RoyalBlue;", BgRoyalBlue(), t) }

func TestFgSaddleBrown(t *testing.T) { test("color:SaddleBrown;", FgSaddleBrown(), t) }
func TestBgSaddleBrown(t *testing.T) { test("background-color:SaddleBrown;", BgSaddleBrown(), t) }

func TestFgSalmon(t *testing.T) { test("color:Salmon;", FgSalmon(), t) }
func TestBgSalmon(t *testing.T) { test("background-color:Salmon;", BgSalmon(), t) }

func TestFgSandyBrown(t *testing.T) { test("color:SandyBrown;", FgSandyBrown(), t) }
func TestBgSandyBrown(t *testing.T) { test("background-color:SandyBrown;", BgSandyBrown(), t) }

func TestFgSeaGreen(t *testing.T) { test("color:SeaGreen;", FgSeaGreen(), t) }
func TestBgSeaGreen(t *testing.T) { test("background-color:SeaGreen;", BgSeaGreen(), t) }

func TestFgSeaShell(t *testing.T) { test("color:SeaShell;", FgSeaShell(), t) }
func TestBgSeaShell(t *testing.T) { test("background-color:SeaShell;", BgSeaShell(), t) }

func TestFgSienna(t *testing.T) { test("color:Sienna;", FgSienna(), t) }
func TestBgSienna(t *testing.T) { test("background-color:Sienna;", BgSienna(), t) }

func TestFgSilver(t *testing.T) { test("color:Silver;", FgSilver(), t) }
func TestBgSilver(t *testing.T) { test("background-color:Silver;", BgSilver(), t) }

func TestFgSkyBlue(t *testing.T) { test("color:SkyBlue;", FgSkyBlue(), t) }
func TestBgSkyBlue(t *testing.T) { test("background-color:SkyBlue;", BgSkyBlue(), t) }

func TestFgSlateBlue(t *testing.T) { test("color:SlateBlue;", FgSlateBlue(), t) }
func TestBgSlateBlue(t *testing.T) { test("background-color:SlateBlue;", BgSlateBlue(), t) }

func TestFgSlateGray(t *testing.T) { test("color:SlateGray;", FgSlateGray(), t) }
func TestBgSlateGray(t *testing.T) { test("background-color:SlateGray;", BgSlateGray(), t) }

func TestFgSlateGrey(t *testing.T) { test("color:SlateGrey;", FgSlateGrey(), t) }
func TestBgSlateGrey(t *testing.T) { test("background-color:SlateGrey;", BgSlateGrey(), t) }

func TestFgSnow(t *testing.T) { test("color:Snow;", FgSnow(), t) }
func TestBgSnow(t *testing.T) { test("background-color:Snow;", BgSnow(), t) }

func TestFgSpringGreen(t *testing.T) { test("color:SpringGreen;", FgSpringGreen(), t) }
func TestBgSpringGreen(t *testing.T) { test("background-color:SpringGreen;", BgSpringGreen(), t) }

func TestFgSteelBlue(t *testing.T) { test("color:SteelBlue;", FgSteelBlue(), t) }
func TestBgSteelBlue(t *testing.T) { test("background-color:SteelBlue;", BgSteelBlue(), t) }

func TestFgTan(t *testing.T) { test("color:Tan;", FgTan(), t) }
func TestBgTan(t *testing.T) { test("background-color:Tan;", BgTan(), t) }

func TestFgTeal(t *testing.T) { test("color:Teal;", FgTeal(), t) }
func TestBgTeal(t *testing.T) { test("background-color:Teal;", BgTeal(), t) }

func TestFgThistle(t *testing.T) { test("color:Thistle;", FgThistle(), t) }
func TestBgThistle(t *testing.T) { test("background-color:Thistle;", BgThistle(), t) }

func TestFgTomato(t *testing.T) { test("color:Tomato;", FgTomato(), t) }
func TestBgTomato(t *testing.T) { test("background-color:Tomato;", BgTomato(), t) }

func TestFgTurquoise(t *testing.T) { test("color:Turquoise;", FgTurquoise(), t) }
func TestBgTurquoise(t *testing.T) { test("background-color:Turquoise;", BgTurquoise(), t) }

func TestFgViolet(t *testing.T) { test("color:Violet;", FgViolet(), t) }
func TestBgViolet(t *testing.T) { test("background-color:Violet;", BgViolet(), t) }

func TestFgWheat(t *testing.T) { test("color:Wheat;", FgWheat(), t) }
func TestBgWheat(t *testing.T) { test("background-color:Wheat;", BgWheat(), t) }

func TestFgWhite(t *testing.T) { test("color:White;", FgWhite(), t) }
func TestBgWhite(t *testing.T) { test("background-color:White;", BgWhite(), t) }

func TestFgWhiteSmoke(t *testing.T) { test("color:WhiteSmoke;", FgWhiteSmoke(), t) }
func TestBgWhiteSmoke(t *testing.T) { test("background-color:WhiteSmoke;", BgWhiteSmoke(), t) }

func TestFgYellow(t *testing.T)   { test("color:Yellow;", FgYellow(), t) }
func TestBgYellow(t *testing.T)   { test("background-color:Yellow;", BgYellow(), t) }
func TestMonospace(t *testing.T)  { test("font-family:'Consola', monospace;", Monospace(), t) }
func TestBold(t *testing.T)       { test("font-weight:bold;", Bold(), t) }
func TestItalic(t *testing.T)     { test("font-style: italic;", Italic(), t) }
func TestUnderline(t *testing.T)  { test("text-decoration: underline;", Underline(), t) }
func TestStrikethru(t *testing.T) { test("text-decoration: line-through;", Strikethru(), t) }

func test(str string, cz colorize.IColorize, t *testing.T) {
	exp := "<span style=\"" + str + "\">Kevin</span>"
	if got := cz.Color("Kevin"); got != exp {
		t.Errorf("expected %q got %q", exp, got)
	}
}
