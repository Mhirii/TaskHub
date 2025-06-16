import { writable } from 'svelte/store';

export interface User {
	username: string;
	email: string;
	avatar: string;
}

const defaultUser: User = {
	username: "",
	email: "",
	avatar: "",
};

function createUserStore() {
	const { subscribe, set, update } = writable<User>(defaultUser);

	return {
		subscribe,
		setUser: (user: User) => set(user),
		clearUser: () => set(defaultUser),
		updateUser: (updates: Partial<User>) =>
			update(u => (u ? { ...u, ...updates } : { ...defaultUser, ...updates })),
	};
}

export const userStore = createUserStore();
