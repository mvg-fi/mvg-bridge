export const filterInputEvents = (event: KeyboardEvent): boolean => {
	if (
		event.code.includes('Digit') || 
		event.code.includes('Backspace') ||
		event.code.includes('Period')
	) return true
	return false
}