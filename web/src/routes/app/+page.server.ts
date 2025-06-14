export function load({ cookies }) {
	const access_token = cookies.get('access_token');

	console.log(access_token);

	return {
		access_token,
	};
}

