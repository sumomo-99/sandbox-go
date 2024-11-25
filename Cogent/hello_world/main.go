package main

import (
	"fmt"
	"strconv"
	"time"

	"cogentcore.org/core/colors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
	"cogentcore.org/core/tree"
)

func main() {
	b := core.NewBody("HelloWorld")
	core.NewButton(b).SetText("Hello, World!")
	core.NewButton(b).SetText("Click me!").SetIcon(icons.Add).OnClick(func(e events.Event) {
		core.MessageSnackbar(b, fmt.Sprint("Button clicked at ", e.Pos()))
	})
	core.NewText(b).SetText("Bold text").Styler(func(s *styles.Style) {
		s.Font.Weight = styles.WeightBold
	})
	core.NewButton(b).SetText("Success button").Styler(func(s *styles.Style) {
		s.Background = colors.Scheme.Success.Base
		s.Color = colors.Scheme.Success.On
	})
	core.NewFrame(b).Styler(func(s *styles.Style) {
		s.Min.Set(units.Dp(50))
		s.Background = colors.Scheme.Primary.Base
	})

	count := 0
	text := core.NewText(b).SetText("0")
	core.NewButton(b).SetText("Increment").OnClick(func(e events.Event) {
		count++
		text.SetText(strconv.Itoa(count)).Update()
	})

	on := true
	core.Bind(&on, core.NewSwitch(b)).OnChange(func(e events.Event) {
		core.MessageSnackbar(b, "The switch is now "+strconv.FormatBool(on))
	})

	number := 3
	spinner := core.Bind(&number, core.NewSpinner(b)).SetMin(0)
	buttons := core.NewFrame(b)
	buttons.Maker(func(p *tree.Plan) {
		for i := range number {
			tree.AddAt(p, strconv.Itoa(i), func(w *core.Button) {
				w.SetText(strconv.Itoa(i))
			})
		}
	})
	spinner.OnChange(func(e events.Event) {
		buttons.Update()
	})

	text2 := core.NewText(b)
	text2.Updater(func() {
		text.SetText(time.Now().Format("15:04:05"))
	})
	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			text2.AsyncLock()
			text2.Update()
			text2.AsyncUnlock()
		}
	}()

	b.RunMainWindow()
}
