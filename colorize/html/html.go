package html

import (
	"fmt"

	"github.com/popmedic/go-color/colorize"
)

// FgAliceBlue returns an IColorize object as a `span` with `style` "color: AliceBlue;"
func FgAliceBlue() colorize.IColorize { return newStyle("color:AliceBlue;") }

// BgAliceBlue returns an IColorize object as a `span` with `style` "background-color: AliceBlue;"
func BgAliceBlue() colorize.IColorize { return newStyle("background-color:AliceBlue;") }

// FgAntiqueWhite returns an IColorize object as a `span` with `style` "color: AntiqueWhite;"
func FgAntiqueWhite() colorize.IColorize { return newStyle("color:AntiqueWhite;") }

// BgAntiqueWhite returns an IColorize object as a `span` with `style` "background-color: AntiqueWhite;"
func BgAntiqueWhite() colorize.IColorize { return newStyle("background-color:AntiqueWhite;") }

// FgAqua returns an IColorize object as a `span` with `style` "color: Aqua;"
func FgAqua() colorize.IColorize { return newStyle("color:Aqua;") }

// BgAqua returns an IColorize object as a `span` with `style` "background-color: Aqua;"
func BgAqua() colorize.IColorize { return newStyle("background-color:Aqua;") }

// FgAquamarine returns an IColorize object as a `span` with `style` "color: Aquamarine;"
func FgAquamarine() colorize.IColorize { return newStyle("color:Aquamarine;") }

// BgAquamarine returns an IColorize object as a `span` with `style` "background-color: Aquamarine;"
func BgAquamarine() colorize.IColorize { return newStyle("background-color:Aquamarine;") }

// FgAzure returns an IColorize object as a `span` with `style` "color: Azure;"
func FgAzure() colorize.IColorize { return newStyle("color:Azure;") }

// BgAzure returns an IColorize object as a `span` with `style` "background-color: Azure;"
func BgAzure() colorize.IColorize { return newStyle("background-color:Azure;") }

// FgBeige returns an IColorize object as a `span` with `style` "color: Beige;"
func FgBeige() colorize.IColorize { return newStyle("color:Beige;") }

// BgBeige returns an IColorize object as a `span` with `style` "background-color: Beige;"
func BgBeige() colorize.IColorize { return newStyle("background-color:Beige;") }

// FgBisque returns an IColorize object as a `span` with `style` "color: Bisque;"
func FgBisque() colorize.IColorize { return newStyle("color:Bisque;") }

// BgBisque returns an IColorize object as a `span` with `style` "background-color: Bisque;"
func BgBisque() colorize.IColorize { return newStyle("background-color:Bisque;") }

// FgBlack returns an IColorize object as a `span` with `style` "color: Black;"
func FgBlack() colorize.IColorize { return newStyle("color:Black;") }

// BgBlack returns an IColorize object as a `span` with `style` "background-color: Black;"
func BgBlack() colorize.IColorize { return newStyle("background-color:Black;") }

// FgBlanchedAlmond returns an IColorize object as a `span` with `style` "color: BlanchedAlmond;"
func FgBlanchedAlmond() colorize.IColorize { return newStyle("color:BlanchedAlmond;") }

// BgBlanchedAlmond returns an IColorize object as a `span` with `style` "background-color: BlanchedAlmond;"
func BgBlanchedAlmond() colorize.IColorize { return newStyle("background-color:BlanchedAlmond;") }

// FgBlue returns an IColorize object as a `span` with `style` "color: Blue;"
func FgBlue() colorize.IColorize { return newStyle("color:Blue;") }

// BgBlue returns an IColorize object as a `span` with `style` "background-color: Blue;"
func BgBlue() colorize.IColorize { return newStyle("background-color:Blue;") }

// FgBlueViolet returns an IColorize object as a `span` with `style` "color: BlueViolet;"
func FgBlueViolet() colorize.IColorize { return newStyle("color:BlueViolet;") }

// BgBlueViolet returns an IColorize object as a `span` with `style` "background-color: BlueViolet;"
func BgBlueViolet() colorize.IColorize { return newStyle("background-color:BlueViolet;") }

// FgBrown returns an IColorize object as a `span` with `style` "color: Brown;"
func FgBrown() colorize.IColorize { return newStyle("color:Brown;") }

// BgBrown returns an IColorize object as a `span` with `style` "background-color: Brown;"
func BgBrown() colorize.IColorize { return newStyle("background-color:Brown;") }

// FgBurlyWood returns an IColorize object as a `span` with `style` "color: BurlyWood;"
func FgBurlyWood() colorize.IColorize { return newStyle("color:BurlyWood;") }

// BgBurlyWood returns an IColorize object as a `span` with `style` "background-color: BurlyWood;"
func BgBurlyWood() colorize.IColorize { return newStyle("background-color:BurlyWood;") }

// FgCadetBlue returns an IColorize object as a `span` with `style` "color: CadetBlue;"
func FgCadetBlue() colorize.IColorize { return newStyle("color:CadetBlue;") }

// BgCadetBlue returns an IColorize object as a `span` with `style` "background-color: CadetBlue;"
func BgCadetBlue() colorize.IColorize { return newStyle("background-color:CadetBlue;") }

// FgChartreuse returns an IColorize object as a `span` with `style` "color: Chartreuse;"
func FgChartreuse() colorize.IColorize { return newStyle("color:Chartreuse;") }

// BgChartreuse returns an IColorize object as a `span` with `style` "background-color: Chartreuse;"
func BgChartreuse() colorize.IColorize { return newStyle("background-color:Chartreuse;") }

// FgChocolate returns an IColorize object as a `span` with `style` "color: Chocolate;"
func FgChocolate() colorize.IColorize { return newStyle("color:Chocolate;") }

// BgChocolate returns an IColorize object as a `span` with `style` "background-color: Chocolate;"
func BgChocolate() colorize.IColorize { return newStyle("background-color:Chocolate;") }

// FgCoral returns an IColorize object as a `span` with `style` "color: Coral;"
func FgCoral() colorize.IColorize { return newStyle("color:Coral;") }

// BgCoral returns an IColorize object as a `span` with `style` "background-color: Coral;"
func BgCoral() colorize.IColorize { return newStyle("background-color:Coral;") }

// FgCornflowerBlue returns an IColorize object as a `span` with `style` "color: CornflowerBlue;"
func FgCornflowerBlue() colorize.IColorize { return newStyle("color:CornflowerBlue;") }

// BgCornflowerBlue returns an IColorize object as a `span` with `style` "background-color: CornflowerBlue;"
func BgCornflowerBlue() colorize.IColorize { return newStyle("background-color:CornflowerBlue;") }

// FgCornsilk returns an IColorize object as a `span` with `style` "color: Cornsilk;"
func FgCornsilk() colorize.IColorize { return newStyle("color:Cornsilk;") }

// BgCornsilk returns an IColorize object as a `span` with `style` "background-color: Cornsilk;"
func BgCornsilk() colorize.IColorize { return newStyle("background-color:Cornsilk;") }

// FgCrimson returns an IColorize object as a `span` with `style` "color: Crimson;"
func FgCrimson() colorize.IColorize { return newStyle("color:Crimson;") }

// BgCrimson returns an IColorize object as a `span` with `style` "background-color: Crimson;"
func BgCrimson() colorize.IColorize { return newStyle("background-color:Crimson;") }

// FgCyan returns an IColorize object as a `span` with `style` "color: Cyan;"
func FgCyan() colorize.IColorize { return newStyle("color:Cyan;") }

// BgCyan returns an IColorize object as a `span` with `style` "background-color: Cyan;"
func BgCyan() colorize.IColorize { return newStyle("background-color:Cyan;") }

// FgDarkBlue returns an IColorize object as a `span` with `style` "color: DarkBlue;"
func FgDarkBlue() colorize.IColorize { return newStyle("color:DarkBlue;") }

// BgDarkBlue returns an IColorize object as a `span` with `style` "background-color: DarkBlue;"
func BgDarkBlue() colorize.IColorize { return newStyle("background-color:DarkBlue;") }

// FgDarkCyan returns an IColorize object as a `span` with `style` "color: DarkCyan;"
func FgDarkCyan() colorize.IColorize { return newStyle("color:DarkCyan;") }

// BgDarkCyan returns an IColorize object as a `span` with `style` "background-color: DarkCyan;"
func BgDarkCyan() colorize.IColorize { return newStyle("background-color:DarkCyan;") }

// FgDarkGoldenRod returns an IColorize object as a `span` with `style` "color: DarkGoldenRod;"
func FgDarkGoldenRod() colorize.IColorize { return newStyle("color:DarkGoldenRod;") }

// BgDarkGoldenRod returns an IColorize object as a `span` with `style` "background-color: DarkGoldenRod;"
func BgDarkGoldenRod() colorize.IColorize { return newStyle("background-color:DarkGoldenRod;") }

// FgDarkGray returns an IColorize object as a `span` with `style` "color: DarkGray;"
func FgDarkGray() colorize.IColorize { return newStyle("color:DarkGray;") }

// BgDarkGray returns an IColorize object as a `span` with `style` "background-color: DarkGray;"
func BgDarkGray() colorize.IColorize { return newStyle("background-color:DarkGray;") }

// FgDarkGrey returns an IColorize object as a `span` with `style` "color: DarkGrey;"
func FgDarkGrey() colorize.IColorize { return newStyle("color:DarkGrey;") }

// BgDarkGrey returns an IColorize object as a `span` with `style` "background-color: DarkGrey;"
func BgDarkGrey() colorize.IColorize { return newStyle("background-color:DarkGrey;") }

// FgDarkGreen returns an IColorize object as a `span` with `style` "color: DarkGreen;"
func FgDarkGreen() colorize.IColorize { return newStyle("color:DarkGreen;") }

// BgDarkGreen returns an IColorize object as a `span` with `style` "background-color: DarkGreen;"
func BgDarkGreen() colorize.IColorize { return newStyle("background-color:DarkGreen;") }

// FgDarkKhaki returns an IColorize object as a `span` with `style` "color: DarkKhaki;"
func FgDarkKhaki() colorize.IColorize { return newStyle("color:DarkKhaki;") }

// BgDarkKhaki returns an IColorize object as a `span` with `style` "background-color: DarkKhaki;"
func BgDarkKhaki() colorize.IColorize { return newStyle("background-color:DarkKhaki;") }

// FgDarkMagenta returns an IColorize object as a `span` with `style` "color: DarkMagenta;"
func FgDarkMagenta() colorize.IColorize { return newStyle("color:DarkMagenta;") }

// BgDarkMagenta returns an IColorize object as a `span` with `style` "background-color: DarkMagenta;"
func BgDarkMagenta() colorize.IColorize { return newStyle("background-color:DarkMagenta;") }

// FgDarkOliveGreen returns an IColorize object as a `span` with `style` "color: DarkOliveGreen;"
func FgDarkOliveGreen() colorize.IColorize { return newStyle("color:DarkOliveGreen;") }

// BgDarkOliveGreen returns an IColorize object as a `span` with `style` "background-color: DarkOliveGreen;"
func BgDarkOliveGreen() colorize.IColorize { return newStyle("background-color:DarkOliveGreen;") }

// FgDarkOrange returns an IColorize object as a `span` with `style` "color: DarkOrange;"
func FgDarkOrange() colorize.IColorize { return newStyle("color:DarkOrange;") }

// BgDarkOrange returns an IColorize object as a `span` with `style` "background-color: DarkOrange;"
func BgDarkOrange() colorize.IColorize { return newStyle("background-color:DarkOrange;") }

// FgDarkOrchid returns an IColorize object as a `span` with `style` "color: DarkOrchid;"
func FgDarkOrchid() colorize.IColorize { return newStyle("color:DarkOrchid;") }

// BgDarkOrchid returns an IColorize object as a `span` with `style` "background-color: DarkOrchid;"
func BgDarkOrchid() colorize.IColorize { return newStyle("background-color:DarkOrchid;") }

// FgDarkRed returns an IColorize object as a `span` with `style` "color: DarkRed;"
func FgDarkRed() colorize.IColorize { return newStyle("color:DarkRed;") }

// BgDarkRed returns an IColorize object as a `span` with `style` "background-color: DarkRed;"
func BgDarkRed() colorize.IColorize { return newStyle("background-color:DarkRed;") }

// FgDarkSalmon returns an IColorize object as a `span` with `style` "color: DarkSalmon;"
func FgDarkSalmon() colorize.IColorize { return newStyle("color:DarkSalmon;") }

// BgDarkSalmon returns an IColorize object as a `span` with `style` "background-color: DarkSalmon;"
func BgDarkSalmon() colorize.IColorize { return newStyle("background-color:DarkSalmon;") }

// FgDarkSeaGreen returns an IColorize object as a `span` with `style` "color: DarkSeaGreen;"
func FgDarkSeaGreen() colorize.IColorize { return newStyle("color:DarkSeaGreen;") }

// BgDarkSeaGreen returns an IColorize object as a `span` with `style` "background-color: DarkSeaGreen;"
func BgDarkSeaGreen() colorize.IColorize { return newStyle("background-color:DarkSeaGreen;") }

// FgDarkSlateBlue returns an IColorize object as a `span` with `style` "color: DarkSlateBlue;"
func FgDarkSlateBlue() colorize.IColorize { return newStyle("color:DarkSlateBlue;") }

// BgDarkSlateBlue returns an IColorize object as a `span` with `style` "background-color: DarkSlateBlue;"
func BgDarkSlateBlue() colorize.IColorize { return newStyle("background-color:DarkSlateBlue;") }

// FgDarkSlateGray returns an IColorize object as a `span` with `style` "color: DarkSlateGray;"
func FgDarkSlateGray() colorize.IColorize { return newStyle("color:DarkSlateGray;") }

// BgDarkSlateGray returns an IColorize object as a `span` with `style` "background-color: DarkSlateGray;"
func BgDarkSlateGray() colorize.IColorize { return newStyle("background-color:DarkSlateGray;") }

// FgDarkSlateGrey returns an IColorize object as a `span` with `style` "color: DarkSlateGrey;"
func FgDarkSlateGrey() colorize.IColorize { return newStyle("color:DarkSlateGrey;") }

// BgDarkSlateGrey returns an IColorize object as a `span` with `style` "background-color: DarkSlateGrey;"
func BgDarkSlateGrey() colorize.IColorize { return newStyle("background-color:DarkSlateGrey;") }

// FgDarkTurquoise returns an IColorize object as a `span` with `style` "color: DarkTurquoise;"
func FgDarkTurquoise() colorize.IColorize { return newStyle("color:DarkTurquoise;") }

// BgDarkTurquoise returns an IColorize object as a `span` with `style` "background-color: DarkTurquoise;"
func BgDarkTurquoise() colorize.IColorize { return newStyle("background-color:DarkTurquoise;") }

// FgDarkViolet returns an IColorize object as a `span` with `style` "color: DarkViolet;"
func FgDarkViolet() colorize.IColorize { return newStyle("color:DarkViolet;") }

// BgDarkViolet returns an IColorize object as a `span` with `style` "background-color: DarkViolet;"
func BgDarkViolet() colorize.IColorize { return newStyle("background-color:DarkViolet;") }

// FgDeepPink returns an IColorize object as a `span` with `style` "color: DeepPink;"
func FgDeepPink() colorize.IColorize { return newStyle("color:DeepPink;") }

// BgDeepPink returns an IColorize object as a `span` with `style` "background-color: DeepPink;"
func BgDeepPink() colorize.IColorize { return newStyle("background-color:DeepPink;") }

// FgDeepSkyBlue returns an IColorize object as a `span` with `style` "color: DeepSkyBlue;"
func FgDeepSkyBlue() colorize.IColorize { return newStyle("color:DeepSkyBlue;") }

// BgDeepSkyBlue returns an IColorize object as a `span` with `style` "background-color: DeepSkyBlue;"
func BgDeepSkyBlue() colorize.IColorize { return newStyle("background-color:DeepSkyBlue;") }

// FgDimGray returns an IColorize object as a `span` with `style` "color: DimGray;"
func FgDimGray() colorize.IColorize { return newStyle("color:DimGray;") }

// BgDimGray returns an IColorize object as a `span` with `style` "background-color: DimGray;"
func BgDimGray() colorize.IColorize { return newStyle("background-color:DimGray;") }

// FgDimGrey returns an IColorize object as a `span` with `style` "color: DimGrey;"
func FgDimGrey() colorize.IColorize { return newStyle("color:DimGrey;") }

// BgDimGrey returns an IColorize object as a `span` with `style` "background-color: DimGrey;"
func BgDimGrey() colorize.IColorize { return newStyle("background-color:DimGrey;") }

// FgDodgerBlue returns an IColorize object as a `span` with `style` "color: DodgerBlue;"
func FgDodgerBlue() colorize.IColorize { return newStyle("color:DodgerBlue;") }

// BgDodgerBlue returns an IColorize object as a `span` with `style` "background-color: DodgerBlue;"
func BgDodgerBlue() colorize.IColorize { return newStyle("background-color:DodgerBlue;") }

// FgFireBrick returns an IColorize object as a `span` with `style` "color: FireBrick;"
func FgFireBrick() colorize.IColorize { return newStyle("color:FireBrick;") }

// BgFireBrick returns an IColorize object as a `span` with `style` "background-color: FireBrick;"
func BgFireBrick() colorize.IColorize { return newStyle("background-color:FireBrick;") }

// FgFloralWhite returns an IColorize object as a `span` with `style` "color: FloralWhite;"
func FgFloralWhite() colorize.IColorize { return newStyle("color:FloralWhite;") }

// BgFloralWhite returns an IColorize object as a `span` with `style` "background-color: FloralWhite;"
func BgFloralWhite() colorize.IColorize { return newStyle("background-color:FloralWhite;") }

// FgForestGreen returns an IColorize object as a `span` with `style` "color: ForestGreen;"
func FgForestGreen() colorize.IColorize { return newStyle("color:ForestGreen;") }

// BgForestGreen returns an IColorize object as a `span` with `style` "background-color: ForestGreen;"
func BgForestGreen() colorize.IColorize { return newStyle("background-color:ForestGreen;") }

// FgFuchsia returns an IColorize object as a `span` with `style` "color: Fuchsia;"
func FgFuchsia() colorize.IColorize { return newStyle("color:Fuchsia;") }

// BgFuchsia returns an IColorize object as a `span` with `style` "background-color: Fuchsia;"
func BgFuchsia() colorize.IColorize { return newStyle("background-color:Fuchsia;") }

// FgGainsboro returns an IColorize object as a `span` with `style` "color: Gainsboro;"
func FgGainsboro() colorize.IColorize { return newStyle("color:Gainsboro;") }

// BgGainsboro returns an IColorize object as a `span` with `style` "background-color: Gainsboro;"
func BgGainsboro() colorize.IColorize { return newStyle("background-color:Gainsboro;") }

// FgGhostWhite returns an IColorize object as a `span` with `style` "color: GhostWhite;"
func FgGhostWhite() colorize.IColorize { return newStyle("color:GhostWhite;") }

// BgGhostWhite returns an IColorize object as a `span` with `style` "background-color: GhostWhite;"
func BgGhostWhite() colorize.IColorize { return newStyle("background-color:GhostWhite;") }

// FgGold returns an IColorize object as a `span` with `style` "color: Gold;"
func FgGold() colorize.IColorize { return newStyle("color:Gold;") }

// BgGold returns an IColorize object as a `span` with `style` "background-color: Gold;"
func BgGold() colorize.IColorize { return newStyle("background-color:Gold;") }

// FgGoldenRod returns an IColorize object as a `span` with `style` "color: GoldenRod;"
func FgGoldenRod() colorize.IColorize { return newStyle("color:GoldenRod;") }

// BgGoldenRod returns an IColorize object as a `span` with `style` "background-color: GoldenRod;"
func BgGoldenRod() colorize.IColorize { return newStyle("background-color:GoldenRod;") }

// FgGray returns an IColorize object as a `span` with `style` "color: Gray;"
func FgGray() colorize.IColorize { return newStyle("color:Gray;") }

// BgGray returns an IColorize object as a `span` with `style` "background-color: Gray;"
func BgGray() colorize.IColorize { return newStyle("background-color:Gray;") }

// FgGrey returns an IColorize object as a `span` with `style` "color: Grey;"
func FgGrey() colorize.IColorize { return newStyle("color:Grey;") }

// BgGrey returns an IColorize object as a `span` with `style` "background-color: Grey;"
func BgGrey() colorize.IColorize { return newStyle("background-color:Grey;") }

// FgGreen returns an IColorize object as a `span` with `style` "color: Green;"
func FgGreen() colorize.IColorize { return newStyle("color:Green;") }

// BgGreen returns an IColorize object as a `span` with `style` "background-color: Green;"
func BgGreen() colorize.IColorize { return newStyle("background-color:Green;") }

// FgGreenYellow returns an IColorize object as a `span` with `style` "color: GreenYellow;"
func FgGreenYellow() colorize.IColorize { return newStyle("color:GreenYellow;") }

// BgGreenYellow returns an IColorize object as a `span` with `style` "background-color: GreenYellow;"
func BgGreenYellow() colorize.IColorize { return newStyle("background-color:GreenYellow;") }

// FgHoneyDew returns an IColorize object as a `span` with `style` "color: HoneyDew;"
func FgHoneyDew() colorize.IColorize { return newStyle("color:HoneyDew;") }

// BgHoneyDew returns an IColorize object as a `span` with `style` "background-color: HoneyDew;"
func BgHoneyDew() colorize.IColorize { return newStyle("background-color:HoneyDew;") }

// FgHotPink returns an IColorize object as a `span` with `style` "color: HotPink;"
func FgHotPink() colorize.IColorize { return newStyle("color:HotPink;") }

// BgHotPink returns an IColorize object as a `span` with `style` "background-color: HotPink;"
func BgHotPink() colorize.IColorize { return newStyle("background-color:HotPink;") }

// FgIndianRed  returns an IColorize object as a `span` with `style` "color: IndianRed ;"
func FgIndianRed() colorize.IColorize { return newStyle("color:IndianRed ;") }

// BgIndianRed  returns an IColorize object as a `span` with `style` "background-color: IndianRed ;"
func BgIndianRed() colorize.IColorize { return newStyle("background-color:IndianRed ;") }

// FgIndigo  returns an IColorize object as a `span` with `style` "color: Indigo ;"
func FgIndigo() colorize.IColorize { return newStyle("color:Indigo ;") }

// BgIndigo  returns an IColorize object as a `span` with `style` "background-color: Indigo ;"
func BgIndigo() colorize.IColorize { return newStyle("background-color:Indigo ;") }

// FgIvory returns an IColorize object as a `span` with `style` "color: Ivory;"
func FgIvory() colorize.IColorize { return newStyle("color:Ivory;") }

// BgIvory returns an IColorize object as a `span` with `style` "background-color: Ivory;"
func BgIvory() colorize.IColorize { return newStyle("background-color:Ivory;") }

// FgKhaki returns an IColorize object as a `span` with `style` "color: Khaki;"
func FgKhaki() colorize.IColorize { return newStyle("color:Khaki;") }

// BgKhaki returns an IColorize object as a `span` with `style` "background-color: Khaki;"
func BgKhaki() colorize.IColorize { return newStyle("background-color:Khaki;") }

// FgLavender returns an IColorize object as a `span` with `style` "color: Lavender;"
func FgLavender() colorize.IColorize { return newStyle("color:Lavender;") }

// BgLavender returns an IColorize object as a `span` with `style` "background-color: Lavender;"
func BgLavender() colorize.IColorize { return newStyle("background-color:Lavender;") }

// FgLavenderBlush returns an IColorize object as a `span` with `style` "color: LavenderBlush;"
func FgLavenderBlush() colorize.IColorize { return newStyle("color:LavenderBlush;") }

// BgLavenderBlush returns an IColorize object as a `span` with `style` "background-color: LavenderBlush;"
func BgLavenderBlush() colorize.IColorize { return newStyle("background-color:LavenderBlush;") }

// FgLawnGreen returns an IColorize object as a `span` with `style` "color: LawnGreen;"
func FgLawnGreen() colorize.IColorize { return newStyle("color:LawnGreen;") }

// BgLawnGreen returns an IColorize object as a `span` with `style` "background-color: LawnGreen;"
func BgLawnGreen() colorize.IColorize { return newStyle("background-color:LawnGreen;") }

// FgLemonChiffon returns an IColorize object as a `span` with `style` "color: LemonChiffon;"
func FgLemonChiffon() colorize.IColorize { return newStyle("color:LemonChiffon;") }

// BgLemonChiffon returns an IColorize object as a `span` with `style` "background-color: LemonChiffon;"
func BgLemonChiffon() colorize.IColorize { return newStyle("background-color:LemonChiffon;") }

// FgLightBlue returns an IColorize object as a `span` with `style` "color: LightBlue;"
func FgLightBlue() colorize.IColorize { return newStyle("color:LightBlue;") }

// BgLightBlue returns an IColorize object as a `span` with `style` "background-color: LightBlue;"
func BgLightBlue() colorize.IColorize { return newStyle("background-color:LightBlue;") }

// FgLightCoral returns an IColorize object as a `span` with `style` "color: LightCoral;"
func FgLightCoral() colorize.IColorize { return newStyle("color:LightCoral;") }

// BgLightCoral returns an IColorize object as a `span` with `style` "background-color: LightCoral;"
func BgLightCoral() colorize.IColorize { return newStyle("background-color:LightCoral;") }

// FgLightCyan returns an IColorize object as a `span` with `style` "color: LightCyan;"
func FgLightCyan() colorize.IColorize { return newStyle("color:LightCyan;") }

// BgLightCyan returns an IColorize object as a `span` with `style` "background-color: LightCyan;"
func BgLightCyan() colorize.IColorize { return newStyle("background-color:LightCyan;") }

// FgLightGoldenRodYellow returns an IColorize object as a `span` with `style` "color: LightGoldenRodYellow;"
func FgLightGoldenRodYellow() colorize.IColorize { return newStyle("color:LightGoldenRodYellow;") }

// BgLightGoldenRodYellow returns an IColorize object as a `span` with `style` "background-color: LightGoldenRodYellow;"
func BgLightGoldenRodYellow() colorize.IColorize {
	return newStyle("background-color:LightGoldenRodYellow;")
}

// FgLightGray returns an IColorize object as a `span` with `style` "color: LightGray;"
func FgLightGray() colorize.IColorize { return newStyle("color:LightGray;") }

// BgLightGray returns an IColorize object as a `span` with `style` "background-color: LightGray;"
func BgLightGray() colorize.IColorize { return newStyle("background-color:LightGray;") }

// FgLightGrey returns an IColorize object as a `span` with `style` "color: LightGrey;"
func FgLightGrey() colorize.IColorize { return newStyle("color:LightGrey;") }

// BgLightGrey returns an IColorize object as a `span` with `style` "background-color: LightGrey;"
func BgLightGrey() colorize.IColorize { return newStyle("background-color:LightGrey;") }

// FgLightGreen returns an IColorize object as a `span` with `style` "color: LightGreen;"
func FgLightGreen() colorize.IColorize { return newStyle("color:LightGreen;") }

// BgLightGreen returns an IColorize object as a `span` with `style` "background-color: LightGreen;"
func BgLightGreen() colorize.IColorize { return newStyle("background-color:LightGreen;") }

// FgLightPink returns an IColorize object as a `span` with `style` "color: LightPink;"
func FgLightPink() colorize.IColorize { return newStyle("color:LightPink;") }

// BgLightPink returns an IColorize object as a `span` with `style` "background-color: LightPink;"
func BgLightPink() colorize.IColorize { return newStyle("background-color:LightPink;") }

// FgLightSalmon returns an IColorize object as a `span` with `style` "color: LightSalmon;"
func FgLightSalmon() colorize.IColorize { return newStyle("color:LightSalmon;") }

// BgLightSalmon returns an IColorize object as a `span` with `style` "background-color: LightSalmon;"
func BgLightSalmon() colorize.IColorize { return newStyle("background-color:LightSalmon;") }

// FgLightSeaGreen returns an IColorize object as a `span` with `style` "color: LightSeaGreen;"
func FgLightSeaGreen() colorize.IColorize { return newStyle("color:LightSeaGreen;") }

// BgLightSeaGreen returns an IColorize object as a `span` with `style` "background-color: LightSeaGreen;"
func BgLightSeaGreen() colorize.IColorize { return newStyle("background-color:LightSeaGreen;") }

// FgLightSkyBlue returns an IColorize object as a `span` with `style` "color: LightSkyBlue;"
func FgLightSkyBlue() colorize.IColorize { return newStyle("color:LightSkyBlue;") }

// BgLightSkyBlue returns an IColorize object as a `span` with `style` "background-color: LightSkyBlue;"
func BgLightSkyBlue() colorize.IColorize { return newStyle("background-color:LightSkyBlue;") }

// FgLightSlateGray returns an IColorize object as a `span` with `style` "color: LightSlateGray;"
func FgLightSlateGray() colorize.IColorize { return newStyle("color:LightSlateGray;") }

// BgLightSlateGray returns an IColorize object as a `span` with `style` "background-color: LightSlateGray;"
func BgLightSlateGray() colorize.IColorize { return newStyle("background-color:LightSlateGray;") }

// FgLightSlateGrey returns an IColorize object as a `span` with `style` "color: LightSlateGrey;"
func FgLightSlateGrey() colorize.IColorize { return newStyle("color:LightSlateGrey;") }

// BgLightSlateGrey returns an IColorize object as a `span` with `style` "background-color: LightSlateGrey;"
func BgLightSlateGrey() colorize.IColorize { return newStyle("background-color:LightSlateGrey;") }

// FgLightSteelBlue returns an IColorize object as a `span` with `style` "color: LightSteelBlue;"
func FgLightSteelBlue() colorize.IColorize { return newStyle("color:LightSteelBlue;") }

// BgLightSteelBlue returns an IColorize object as a `span` with `style` "background-color: LightSteelBlue;"
func BgLightSteelBlue() colorize.IColorize { return newStyle("background-color:LightSteelBlue;") }

// FgLightYellow returns an IColorize object as a `span` with `style` "color: LightYellow;"
func FgLightYellow() colorize.IColorize { return newStyle("color:LightYellow;") }

// BgLightYellow returns an IColorize object as a `span` with `style` "background-color: LightYellow;"
func BgLightYellow() colorize.IColorize { return newStyle("background-color:LightYellow;") }

// FgLime returns an IColorize object as a `span` with `style` "color: Lime;"
func FgLime() colorize.IColorize { return newStyle("color:Lime;") }

// BgLime returns an IColorize object as a `span` with `style` "background-color: Lime;"
func BgLime() colorize.IColorize { return newStyle("background-color:Lime;") }

// FgLimeGreen returns an IColorize object as a `span` with `style` "color: LimeGreen;"
func FgLimeGreen() colorize.IColorize { return newStyle("color:LimeGreen;") }

// BgLimeGreen returns an IColorize object as a `span` with `style` "background-color: LimeGreen;"
func BgLimeGreen() colorize.IColorize { return newStyle("background-color:LimeGreen;") }

// FgLinen returns an IColorize object as a `span` with `style` "color: Linen;"
func FgLinen() colorize.IColorize { return newStyle("color:Linen;") }

// BgLinen returns an IColorize object as a `span` with `style` "background-color: Linen;"
func BgLinen() colorize.IColorize { return newStyle("background-color:Linen;") }

// FgMagenta returns an IColorize object as a `span` with `style` "color: Magenta;"
func FgMagenta() colorize.IColorize { return newStyle("color:Magenta;") }

// BgMagenta returns an IColorize object as a `span` with `style` "background-color: Magenta;"
func BgMagenta() colorize.IColorize { return newStyle("background-color:Magenta;") }

// FgMaroon returns an IColorize object as a `span` with `style` "color: Maroon;"
func FgMaroon() colorize.IColorize { return newStyle("color:Maroon;") }

// BgMaroon returns an IColorize object as a `span` with `style` "background-color: Maroon;"
func BgMaroon() colorize.IColorize { return newStyle("background-color:Maroon;") }

// FgMediumAquaMarine returns an IColorize object as a `span` with `style` "color: MediumAquaMarine;"
func FgMediumAquaMarine() colorize.IColorize { return newStyle("color:MediumAquaMarine;") }

// BgMediumAquaMarine returns an IColorize object as a `span` with `style` "background-color: MediumAquaMarine;"
func BgMediumAquaMarine() colorize.IColorize { return newStyle("background-color:MediumAquaMarine;") }

// FgMediumBlue returns an IColorize object as a `span` with `style` "color: MediumBlue;"
func FgMediumBlue() colorize.IColorize { return newStyle("color:MediumBlue;") }

// BgMediumBlue returns an IColorize object as a `span` with `style` "background-color: MediumBlue;"
func BgMediumBlue() colorize.IColorize { return newStyle("background-color:MediumBlue;") }

// FgMediumOrchid returns an IColorize object as a `span` with `style` "color: MediumOrchid;"
func FgMediumOrchid() colorize.IColorize { return newStyle("color:MediumOrchid;") }

// BgMediumOrchid returns an IColorize object as a `span` with `style` "background-color: MediumOrchid;"
func BgMediumOrchid() colorize.IColorize { return newStyle("background-color:MediumOrchid;") }

// FgMediumPurple returns an IColorize object as a `span` with `style` "color: MediumPurple;"
func FgMediumPurple() colorize.IColorize { return newStyle("color:MediumPurple;") }

// BgMediumPurple returns an IColorize object as a `span` with `style` "background-color: MediumPurple;"
func BgMediumPurple() colorize.IColorize { return newStyle("background-color:MediumPurple;") }

// FgMediumSeaGreen returns an IColorize object as a `span` with `style` "color: MediumSeaGreen;"
func FgMediumSeaGreen() colorize.IColorize { return newStyle("color:MediumSeaGreen;") }

// BgMediumSeaGreen returns an IColorize object as a `span` with `style` "background-color: MediumSeaGreen;"
func BgMediumSeaGreen() colorize.IColorize { return newStyle("background-color:MediumSeaGreen;") }

// FgMediumSlateBlue returns an IColorize object as a `span` with `style` "color: MediumSlateBlue;"
func FgMediumSlateBlue() colorize.IColorize { return newStyle("color:MediumSlateBlue;") }

// BgMediumSlateBlue returns an IColorize object as a `span` with `style` "background-color: MediumSlateBlue;"
func BgMediumSlateBlue() colorize.IColorize { return newStyle("background-color:MediumSlateBlue;") }

// FgMediumSpringGreen returns an IColorize object as a `span` with `style` "color: MediumSpringGreen;"
func FgMediumSpringGreen() colorize.IColorize { return newStyle("color:MediumSpringGreen;") }

// BgMediumSpringGreen returns an IColorize object as a `span` with `style` "background-color: MediumSpringGreen;"
func BgMediumSpringGreen() colorize.IColorize { return newStyle("background-color:MediumSpringGreen;") }

// FgMediumTurquoise returns an IColorize object as a `span` with `style` "color: MediumTurquoise;"
func FgMediumTurquoise() colorize.IColorize { return newStyle("color:MediumTurquoise;") }

// BgMediumTurquoise returns an IColorize object as a `span` with `style` "background-color: MediumTurquoise;"
func BgMediumTurquoise() colorize.IColorize { return newStyle("background-color:MediumTurquoise;") }

// FgMediumVioletRed returns an IColorize object as a `span` with `style` "color: MediumVioletRed;"
func FgMediumVioletRed() colorize.IColorize { return newStyle("color:MediumVioletRed;") }

// BgMediumVioletRed returns an IColorize object as a `span` with `style` "background-color: MediumVioletRed;"
func BgMediumVioletRed() colorize.IColorize { return newStyle("background-color:MediumVioletRed;") }

// FgMidnightBlue returns an IColorize object as a `span` with `style` "color: MidnightBlue;"
func FgMidnightBlue() colorize.IColorize { return newStyle("color:MidnightBlue;") }

// BgMidnightBlue returns an IColorize object as a `span` with `style` "background-color: MidnightBlue;"
func BgMidnightBlue() colorize.IColorize { return newStyle("background-color:MidnightBlue;") }

// FgMintCream returns an IColorize object as a `span` with `style` "color: MintCream;"
func FgMintCream() colorize.IColorize { return newStyle("color:MintCream;") }

// BgMintCream returns an IColorize object as a `span` with `style` "background-color: MintCream;"
func BgMintCream() colorize.IColorize { return newStyle("background-color:MintCream;") }

// FgMistyRose returns an IColorize object as a `span` with `style` "color: MistyRose;"
func FgMistyRose() colorize.IColorize { return newStyle("color:MistyRose;") }

// BgMistyRose returns an IColorize object as a `span` with `style` "background-color: MistyRose;"
func BgMistyRose() colorize.IColorize { return newStyle("background-color:MistyRose;") }

// FgMoccasin returns an IColorize object as a `span` with `style` "color: Moccasin;"
func FgMoccasin() colorize.IColorize { return newStyle("color:Moccasin;") }

// BgMoccasin returns an IColorize object as a `span` with `style` "background-color: Moccasin;"
func BgMoccasin() colorize.IColorize { return newStyle("background-color:Moccasin;") }

// FgNavajoWhite returns an IColorize object as a `span` with `style` "color: NavajoWhite;"
func FgNavajoWhite() colorize.IColorize { return newStyle("color:NavajoWhite;") }

// BgNavajoWhite returns an IColorize object as a `span` with `style` "background-color: NavajoWhite;"
func BgNavajoWhite() colorize.IColorize { return newStyle("background-color:NavajoWhite;") }

// FgNavy returns an IColorize object as a `span` with `style` "color: Navy;"
func FgNavy() colorize.IColorize { return newStyle("color:Navy;") }

// BgNavy returns an IColorize object as a `span` with `style` "background-color: Navy;"
func BgNavy() colorize.IColorize { return newStyle("background-color:Navy;") }

// FgOldLace returns an IColorize object as a `span` with `style` "color: OldLace;"
func FgOldLace() colorize.IColorize { return newStyle("color:OldLace;") }

// BgOldLace returns an IColorize object as a `span` with `style` "background-color: OldLace;"
func BgOldLace() colorize.IColorize { return newStyle("background-color:OldLace;") }

// FgOlive returns an IColorize object as a `span` with `style` "color: Olive;"
func FgOlive() colorize.IColorize { return newStyle("color:Olive;") }

// BgOlive returns an IColorize object as a `span` with `style` "background-color: Olive;"
func BgOlive() colorize.IColorize { return newStyle("background-color:Olive;") }

// FgOliveDrab returns an IColorize object as a `span` with `style` "color: OliveDrab;"
func FgOliveDrab() colorize.IColorize { return newStyle("color:OliveDrab;") }

// BgOliveDrab returns an IColorize object as a `span` with `style` "background-color: OliveDrab;"
func BgOliveDrab() colorize.IColorize { return newStyle("background-color:OliveDrab;") }

// FgOrange returns an IColorize object as a `span` with `style` "color: Orange;"
func FgOrange() colorize.IColorize { return newStyle("color:Orange;") }

// BgOrange returns an IColorize object as a `span` with `style` "background-color: Orange;"
func BgOrange() colorize.IColorize { return newStyle("background-color:Orange;") }

// FgOrangeRed returns an IColorize object as a `span` with `style` "color: OrangeRed;"
func FgOrangeRed() colorize.IColorize { return newStyle("color:OrangeRed;") }

// BgOrangeRed returns an IColorize object as a `span` with `style` "background-color: OrangeRed;"
func BgOrangeRed() colorize.IColorize { return newStyle("background-color:OrangeRed;") }

// FgOrchid returns an IColorize object as a `span` with `style` "color: Orchid;"
func FgOrchid() colorize.IColorize { return newStyle("color:Orchid;") }

// BgOrchid returns an IColorize object as a `span` with `style` "background-color: Orchid;"
func BgOrchid() colorize.IColorize { return newStyle("background-color:Orchid;") }

// FgPaleGoldenRod returns an IColorize object as a `span` with `style` "color: PaleGoldenRod;"
func FgPaleGoldenRod() colorize.IColorize { return newStyle("color:PaleGoldenRod;") }

// BgPaleGoldenRod returns an IColorize object as a `span` with `style` "background-color: PaleGoldenRod;"
func BgPaleGoldenRod() colorize.IColorize { return newStyle("background-color:PaleGoldenRod;") }

// FgPaleGreen returns an IColorize object as a `span` with `style` "color: PaleGreen;"
func FgPaleGreen() colorize.IColorize { return newStyle("color:PaleGreen;") }

// BgPaleGreen returns an IColorize object as a `span` with `style` "background-color: PaleGreen;"
func BgPaleGreen() colorize.IColorize { return newStyle("background-color:PaleGreen;") }

// FgPaleTurquoise returns an IColorize object as a `span` with `style` "color: PaleTurquoise;"
func FgPaleTurquoise() colorize.IColorize { return newStyle("color:PaleTurquoise;") }

// BgPaleTurquoise returns an IColorize object as a `span` with `style` "background-color: PaleTurquoise;"
func BgPaleTurquoise() colorize.IColorize { return newStyle("background-color:PaleTurquoise;") }

// FgPaleVioletRed returns an IColorize object as a `span` with `style` "color: PaleVioletRed;"
func FgPaleVioletRed() colorize.IColorize { return newStyle("color:PaleVioletRed;") }

// BgPaleVioletRed returns an IColorize object as a `span` with `style` "background-color: PaleVioletRed;"
func BgPaleVioletRed() colorize.IColorize { return newStyle("background-color:PaleVioletRed;") }

// FgPapayaWhip returns an IColorize object as a `span` with `style` "color: PapayaWhip;"
func FgPapayaWhip() colorize.IColorize { return newStyle("color:PapayaWhip;") }

// BgPapayaWhip returns an IColorize object as a `span` with `style` "background-color: PapayaWhip;"
func BgPapayaWhip() colorize.IColorize { return newStyle("background-color:PapayaWhip;") }

// FgPeachPuff returns an IColorize object as a `span` with `style` "color: PeachPuff;"
func FgPeachPuff() colorize.IColorize { return newStyle("color:PeachPuff;") }

// BgPeachPuff returns an IColorize object as a `span` with `style` "background-color: PeachPuff;"
func BgPeachPuff() colorize.IColorize { return newStyle("background-color:PeachPuff;") }

// FgPeru returns an IColorize object as a `span` with `style` "color: Peru;"
func FgPeru() colorize.IColorize { return newStyle("color:Peru;") }

// BgPeru returns an IColorize object as a `span` with `style` "background-color: Peru;"
func BgPeru() colorize.IColorize { return newStyle("background-color:Peru;") }

// FgPink returns an IColorize object as a `span` with `style` "color: Pink;"
func FgPink() colorize.IColorize { return newStyle("color:Pink;") }

// BgPink returns an IColorize object as a `span` with `style` "background-color: Pink;"
func BgPink() colorize.IColorize { return newStyle("background-color:Pink;") }

// FgPlum returns an IColorize object as a `span` with `style` "color: Plum;"
func FgPlum() colorize.IColorize { return newStyle("color:Plum;") }

// BgPlum returns an IColorize object as a `span` with `style` "background-color: Plum;"
func BgPlum() colorize.IColorize { return newStyle("background-color:Plum;") }

// FgPowderBlue returns an IColorize object as a `span` with `style` "color: PowderBlue;"
func FgPowderBlue() colorize.IColorize { return newStyle("color:PowderBlue;") }

// BgPowderBlue returns an IColorize object as a `span` with `style` "background-color: PowderBlue;"
func BgPowderBlue() colorize.IColorize { return newStyle("background-color:PowderBlue;") }

// FgPurple returns an IColorize object as a `span` with `style` "color: Purple;"
func FgPurple() colorize.IColorize { return newStyle("color:Purple;") }

// BgPurple returns an IColorize object as a `span` with `style` "background-color: Purple;"
func BgPurple() colorize.IColorize { return newStyle("background-color:Purple;") }

// FgRebeccaPurple returns an IColorize object as a `span` with `style` "color: RebeccaPurple;"
func FgRebeccaPurple() colorize.IColorize { return newStyle("color:RebeccaPurple;") }

// BgRebeccaPurple returns an IColorize object as a `span` with `style` "background-color: RebeccaPurple;"
func BgRebeccaPurple() colorize.IColorize { return newStyle("background-color:RebeccaPurple;") }

// FgRed returns an IColorize object as a `span` with `style` "color: Red;"
func FgRed() colorize.IColorize { return newStyle("color:Red;") }

// BgRed returns an IColorize object as a `span` with `style` "background-color: Red;"
func BgRed() colorize.IColorize { return newStyle("background-color:Red;") }

// FgRosyBrown returns an IColorize object as a `span` with `style` "color: RosyBrown;"
func FgRosyBrown() colorize.IColorize { return newStyle("color:RosyBrown;") }

// BgRosyBrown returns an IColorize object as a `span` with `style` "background-color: RosyBrown;"
func BgRosyBrown() colorize.IColorize { return newStyle("background-color:RosyBrown;") }

// FgRoyalBlue returns an IColorize object as a `span` with `style` "color: RoyalBlue;"
func FgRoyalBlue() colorize.IColorize { return newStyle("color:RoyalBlue;") }

// BgRoyalBlue returns an IColorize object as a `span` with `style` "background-color: RoyalBlue;"
func BgRoyalBlue() colorize.IColorize { return newStyle("background-color:RoyalBlue;") }

// FgSaddleBrown returns an IColorize object as a `span` with `style` "color: SaddleBrown;"
func FgSaddleBrown() colorize.IColorize { return newStyle("color:SaddleBrown;") }

// BgSaddleBrown returns an IColorize object as a `span` with `style` "background-color: SaddleBrown;"
func BgSaddleBrown() colorize.IColorize { return newStyle("background-color:SaddleBrown;") }

// FgSalmon returns an IColorize object as a `span` with `style` "color: Salmon;"
func FgSalmon() colorize.IColorize { return newStyle("color:Salmon;") }

// BgSalmon returns an IColorize object as a `span` with `style` "background-color: Salmon;"
func BgSalmon() colorize.IColorize { return newStyle("background-color:Salmon;") }

// FgSandyBrown returns an IColorize object as a `span` with `style` "color: SandyBrown;"
func FgSandyBrown() colorize.IColorize { return newStyle("color:SandyBrown;") }

// BgSandyBrown returns an IColorize object as a `span` with `style` "background-color: SandyBrown;"
func BgSandyBrown() colorize.IColorize { return newStyle("background-color:SandyBrown;") }

// FgSeaGreen returns an IColorize object as a `span` with `style` "color: SeaGreen;"
func FgSeaGreen() colorize.IColorize { return newStyle("color:SeaGreen;") }

// BgSeaGreen returns an IColorize object as a `span` with `style` "background-color: SeaGreen;"
func BgSeaGreen() colorize.IColorize { return newStyle("background-color:SeaGreen;") }

// FgSeaShell returns an IColorize object as a `span` with `style` "color: SeaShell;"
func FgSeaShell() colorize.IColorize { return newStyle("color:SeaShell;") }

// BgSeaShell returns an IColorize object as a `span` with `style` "background-color: SeaShell;"
func BgSeaShell() colorize.IColorize { return newStyle("background-color:SeaShell;") }

// FgSienna returns an IColorize object as a `span` with `style` "color: Sienna;"
func FgSienna() colorize.IColorize { return newStyle("color:Sienna;") }

// BgSienna returns an IColorize object as a `span` with `style` "background-color: Sienna;"
func BgSienna() colorize.IColorize { return newStyle("background-color:Sienna;") }

// FgSilver returns an IColorize object as a `span` with `style` "color: Silver;"
func FgSilver() colorize.IColorize { return newStyle("color:Silver;") }

// BgSilver returns an IColorize object as a `span` with `style` "background-color: Silver;"
func BgSilver() colorize.IColorize { return newStyle("background-color:Silver;") }

// FgSkyBlue returns an IColorize object as a `span` with `style` "color: SkyBlue;"
func FgSkyBlue() colorize.IColorize { return newStyle("color:SkyBlue;") }

// BgSkyBlue returns an IColorize object as a `span` with `style` "background-color: SkyBlue;"
func BgSkyBlue() colorize.IColorize { return newStyle("background-color:SkyBlue;") }

// FgSlateBlue returns an IColorize object as a `span` with `style` "color: SlateBlue;"
func FgSlateBlue() colorize.IColorize { return newStyle("color:SlateBlue;") }

// BgSlateBlue returns an IColorize object as a `span` with `style` "background-color: SlateBlue;"
func BgSlateBlue() colorize.IColorize { return newStyle("background-color:SlateBlue;") }

// FgSlateGray returns an IColorize object as a `span` with `style` "color: SlateGray;"
func FgSlateGray() colorize.IColorize { return newStyle("color:SlateGray;") }

// BgSlateGray returns an IColorize object as a `span` with `style` "background-color: SlateGray;"
func BgSlateGray() colorize.IColorize { return newStyle("background-color:SlateGray;") }

// FgSlateGrey returns an IColorize object as a `span` with `style` "color: SlateGrey;"
func FgSlateGrey() colorize.IColorize { return newStyle("color:SlateGrey;") }

// BgSlateGrey returns an IColorize object as a `span` with `style` "background-color: SlateGrey;"
func BgSlateGrey() colorize.IColorize { return newStyle("background-color:SlateGrey;") }

// FgSnow returns an IColorize object as a `span` with `style` "color: Snow;"
func FgSnow() colorize.IColorize { return newStyle("color:Snow;") }

// BgSnow returns an IColorize object as a `span` with `style` "background-color: Snow;"
func BgSnow() colorize.IColorize { return newStyle("background-color:Snow;") }

// FgSpringGreen returns an IColorize object as a `span` with `style` "color: SpringGreen;"
func FgSpringGreen() colorize.IColorize { return newStyle("color:SpringGreen;") }

// BgSpringGreen returns an IColorize object as a `span` with `style` "background-color: SpringGreen;"
func BgSpringGreen() colorize.IColorize { return newStyle("background-color:SpringGreen;") }

// FgSteelBlue returns an IColorize object as a `span` with `style` "color: SteelBlue;"
func FgSteelBlue() colorize.IColorize { return newStyle("color:SteelBlue;") }

// BgSteelBlue returns an IColorize object as a `span` with `style` "background-color: SteelBlue;"
func BgSteelBlue() colorize.IColorize { return newStyle("background-color:SteelBlue;") }

// FgTan returns an IColorize object as a `span` with `style` "color: Tan;"
func FgTan() colorize.IColorize { return newStyle("color:Tan;") }

// BgTan returns an IColorize object as a `span` with `style` "background-color: Tan;"
func BgTan() colorize.IColorize { return newStyle("background-color:Tan;") }

// FgTeal returns an IColorize object as a `span` with `style` "color: Teal;"
func FgTeal() colorize.IColorize { return newStyle("color:Teal;") }

// BgTeal returns an IColorize object as a `span` with `style` "background-color: Teal;"
func BgTeal() colorize.IColorize { return newStyle("background-color:Teal;") }

// FgThistle returns an IColorize object as a `span` with `style` "color: Thistle;"
func FgThistle() colorize.IColorize { return newStyle("color:Thistle;") }

// BgThistle returns an IColorize object as a `span` with `style` "background-color: Thistle;"
func BgThistle() colorize.IColorize { return newStyle("background-color:Thistle;") }

// FgTomato returns an IColorize object as a `span` with `style` "color: Tomato;"
func FgTomato() colorize.IColorize { return newStyle("color:Tomato;") }

// BgTomato returns an IColorize object as a `span` with `style` "background-color: Tomato;"
func BgTomato() colorize.IColorize { return newStyle("background-color:Tomato;") }

// FgTurquoise returns an IColorize object as a `span` with `style` "color: Turquoise;"
func FgTurquoise() colorize.IColorize { return newStyle("color:Turquoise;") }

// BgTurquoise returns an IColorize object as a `span` with `style` "background-color: Turquoise;"
func BgTurquoise() colorize.IColorize { return newStyle("background-color:Turquoise;") }

// FgViolet returns an IColorize object as a `span` with `style` "color: Violet;"
func FgViolet() colorize.IColorize { return newStyle("color:Violet;") }

// BgViolet returns an IColorize object as a `span` with `style` "background-color: Violet;"
func BgViolet() colorize.IColorize { return newStyle("background-color:Violet;") }

// FgWheat returns an IColorize object as a `span` with `style` "color: Wheat;"
func FgWheat() colorize.IColorize { return newStyle("color:Wheat;") }

// BgWheat returns an IColorize object as a `span` with `style` "background-color: Wheat;"
func BgWheat() colorize.IColorize { return newStyle("background-color:Wheat;") }

// FgWhite returns an IColorize object as a `span` with `style` "color: White;"
func FgWhite() colorize.IColorize { return newStyle("color:White;") }

// BgWhite returns an IColorize object as a `span` with `style` "background-color: White;"
func BgWhite() colorize.IColorize { return newStyle("background-color:White;") }

// FgWhiteSmoke returns an IColorize object as a `span` with `style` "color: WhiteSmoke;"
func FgWhiteSmoke() colorize.IColorize { return newStyle("color:WhiteSmoke;") }

// BgWhiteSmoke returns an IColorize object as a `span` with `style` "background-color: WhiteSmoke;"
func BgWhiteSmoke() colorize.IColorize { return newStyle("background-color:WhiteSmoke;") }

// FgYellow returns an IColorize object as a `span` with `style` "color: Yellow;"
func FgYellow() colorize.IColorize { return newStyle("color:Yellow;") }

// BgYellow returns an IColorize object as a `span` with `style` "background-color: Yellow;"
func BgYellow() colorize.IColorize { return newStyle("background-color:Yellow;") }

// Monospace returns an IColorize object as a `span` with `style` "font-family:'Consola', monospace;"
func Monospace() colorize.IColorize { return newStyle("font-family:'Consola', monospace;") }

// Bold returns an IColorize object as a `span` with `style` "font-weight:bold;"
func Bold() colorize.IColorize { return newStyle("font-weight:bold;") }

// Italic returns an IColorize object as a `span` with `style` "font-style: italic;"
func Italic() colorize.IColorize { return newStyle("font-style: italic;") }

// Underline returns an IColorize object as a `span` with `style` "font-style: italic;"
func Underline() colorize.IColorize { return newStyle("text-decoration: underline;") }

// Strikethru returns an IColorize object as a `span` with `style` "font-style: italic;"
func Strikethru() colorize.IColorize { return newStyle("text-decoration: line-through;") }

type attrs map[string]string
type attrss []attrs

func (a attrs) String() (rtn string) {
	for k, v := range a {
		rtn = fmt.Sprintf("%s %s=%q", rtn, k, v)
	}
	return
}

func (a attrss) String() (rtn string) {
	for _, v := range a {
		rtn = fmt.Sprintf("%s%s", rtn, v.String())
	}
	return
}

func newAttrs(ss ...string) attrs {
	rtn := make(map[string]string)
	if len(ss)%2 != 0 {
		ss = append(ss, "")
	}
	for i := 0; i < len(ss); i = i + 2 {
		rtn[ss[i]] = ss[i+1]
	}
	return rtn
}

func newTag(tag string, a ...attrs) colorize.IColorize {
	return colorize.NewColorize(fmt.Sprintf("<%s%s>", tag, attrss(a).String()), fmt.Sprintf("</%s>", tag))
}

func newSpan(a ...string) colorize.IColorize {
	return newTag("span", newAttrs(a...))
}

func newStyle(s string) colorize.IColorize {
	return newSpan("style", s)
}
