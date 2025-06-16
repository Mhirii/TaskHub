import { USERS_API_URL } from "$env/static/private";
import type { GetUsersResponse } from "$lib/types";
import { superValidate } from "sveltekit-superforms";
import { userSchema } from "./schema";
import { zod } from "sveltekit-superforms/adapters";

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
	const userForm = await superValidate(zod(userSchema))
	const users: GetUsersResponse = await getUsers()
	return {
		users,
		userForm,
	};
}
