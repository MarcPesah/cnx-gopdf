package gopdf

// Margins type.
type Margins struct {
	Left, Top, Right, Bottom float64
}

// SetLeftMargin sets left margin.
func (gp *GoPdf2) SetLeftMargin(margin float64) {
	gp.UnitsToPointsVar(&margin)
	gp.margins.Left = margin
}

// SetTopMargin sets top margin.
func (gp *GoPdf2) SetTopMargin(margin float64) {
	gp.UnitsToPointsVar(&margin)
	gp.margins.Top = margin
}

// SetMargins defines the left, top, right and bottom margins. By default, they equal 1 cm. Call this method to change them.
func (gp *GoPdf2) SetMargins(left, top, right, bottom float64) {
	gp.UnitsToPointsVar(&left, &top, &right, &bottom)
	gp.margins = Margins{left, top, right, bottom}
}

// SetMarginLeft sets the left margin
func (gp *GoPdf2) SetMarginLeft(margin float64) {
	gp.margins.Left = gp.UnitsToPoints(margin)
}

// SetMarginTop sets the top margin
func (gp *GoPdf2) SetMarginTop(margin float64) {
	gp.margins.Top = gp.UnitsToPoints(margin)
}

// SetMarginRight sets the right margin
func (gp *GoPdf2) SetMarginRight(margin float64) {
	gp.margins.Right = gp.UnitsToPoints(margin)
}

// SetMarginBottom set the bottom margin
func (gp *GoPdf2) SetMarginBottom(margin float64) {
	gp.margins.Bottom = gp.UnitsToPoints(margin)
}

// Margins gets the current margins, The margins will be converted back to the documents units. Returned values will be in the following order Left, Top, Right, Bottom
func (gp *GoPdf2) Margins() (float64, float64, float64, float64) {
	return gp.PointsToUnits(gp.margins.Left),
		gp.PointsToUnits(gp.margins.Top),
		gp.PointsToUnits(gp.margins.Right),
		gp.PointsToUnits(gp.margins.Bottom)
}

// MarginLeft returns the left margin.
func (gp *GoPdf2) MarginLeft() float64 {
	return gp.PointsToUnits(gp.margins.Left)
}

// MarginTop returns the top margin.
func (gp *GoPdf2) MarginTop() float64 {
	return gp.PointsToUnits(gp.margins.Top)
}

// MarginRight returns the right margin.
func (gp *GoPdf2) MarginRight() float64 {
	return gp.PointsToUnits(gp.margins.Right)
}

// MarginBottom returns the bottom margin.
func (gp *GoPdf2) MarginBottom() float64 {
	return gp.PointsToUnits(gp.margins.Bottom)
}
