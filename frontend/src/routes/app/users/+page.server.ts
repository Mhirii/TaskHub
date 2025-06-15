import { USERS_API_URL } from "$env/static/private";
import type { GetUsersResponse } from "$lib/types";

export async function load({ cookies }) {
	const access_token = cookies.get('access_token');
	const getUsers = async () => fetch(`${USERS_API_URL}/`, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json',
			'Authorization': `Bearer ${access_token}`,
		}
	}).then(res => {
		if (!res.ok) throw new Error('Failed to fetch users');
		const result = res.json()
		console.log(result)
		return result
	});
	const users: GetUsersResponse = await getUsers()
	return {
		users,
	};
}
