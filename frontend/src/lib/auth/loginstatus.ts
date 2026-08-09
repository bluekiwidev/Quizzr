export default function hasCookie(name: string): boolean {
	return document.cookie.split(';').some((item) => {
		const trimmed = item.trim();
		return trimmed.startsWith(`${name}=`);
	});
}
