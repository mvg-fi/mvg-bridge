export const filterInputEvents = (event: KeyboardEvent): boolean => {
	if (
		event.code.includes('Digit') ||
		event.code.includes('Backspace') ||
		event.code.includes('Period')
	) return true
	return false
}

export const shortenAddress = (input: string): string => {
	if (input.length > 16) {
		return `${input.slice(0, 10)}...${input.slice(-6)}`;
	}
	return input;
}

export const toHex = (num: string | number) => {
	const val = Number(num);
	return '0x' + val.toString(16);
};

export const formatUSMoney = (x: string | number) => {
	return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' }).format(Number(x))
}