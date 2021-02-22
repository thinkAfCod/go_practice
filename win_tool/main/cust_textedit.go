package main

import (
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

type CustTextEdit struct {
	Accessibility      declarative.Accessibility
	Background         declarative.Brush
	ContextMenuItems   []declarative.MenuItem
	DoubleBuffering    bool
	Enabled            declarative.Property
	Font               declarative.Font
	MaxSize            declarative.Size
	MinSize            declarative.Size
	Name               string
	OnBoundsChanged    walk.EventHandler
	OnKeyDown          walk.KeyEventHandler
	OnKeyPress         walk.KeyEventHandler
	OnKeyUp            walk.KeyEventHandler
	OnMouseDown        walk.MouseEventHandler
	OnMouseMove        walk.MouseEventHandler
	OnMouseUp          walk.MouseEventHandler
	OnSizeChanged      walk.EventHandler
	Persistent         bool
	RightToLeftReading bool
	ToolTipText        declarative.Property
	Visible            declarative.Property

	// Widget

	Alignment          declarative.Alignment2D
	AlwaysConsumeSpace bool
	Column             int
	ColumnSpan         int
	GraphicsEffects    []walk.WidgetGraphicsEffect
	Row                int
	RowSpan            int
	StretchFactor      int

	// TextEdit

	AssignTo      **walk.TextEdit
	CompactHeight bool
	HScroll       bool
	MaxLength     int
	OnTextChanged walk.EventHandler
	ReadOnly      declarative.Property
	Text          declarative.Property
	TextAlignment declarative.Alignment1D
	TextColor     walk.Color
	VScroll       bool
	OnDropFiles   walk.DropFilesEventHandler
}

func (te CustTextEdit) Create(builder *declarative.Builder) error {
	var style uint32
	if te.HScroll {
		style |= win.WS_HSCROLL
	}
	if te.VScroll {
		style |= win.WS_VSCROLL
	}

	w, err := walk.NewTextEditWithStyle(builder.Parent(), style)
	if err != nil {
		return err
	}

	if te.AssignTo != nil {
		*te.AssignTo = w
	}

	return builder.InitWidget(te, w, func() error {
		w.SetCompactHeight(te.CompactHeight)
		w.SetTextColor(te.TextColor)

		if err := w.SetTextAlignment(walk.Alignment1D(te.TextAlignment)); err != nil {
			return err
		}

		if te.MaxLength > 0 {
			w.SetMaxLength(te.MaxLength)
		}

		if te.OnTextChanged != nil {
			w.TextChanged().Attach(te.OnTextChanged)
		}

		if te.OnDropFiles != nil {
			w.DropFiles().Attach(te.OnDropFiles)
		}
		return nil
	})
}
