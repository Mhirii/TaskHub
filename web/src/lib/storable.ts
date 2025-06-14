
import { writable, type Writable } from 'svelte/store'

export function storable<T>(name: string, data: T) {
	const store = writable(data);
	const { subscribe, set, update } = store;
	const isBrowser = typeof window !== 'undefined';

	isBrowser &&
		localStorage.storable &&
		set(JSON.parse(localStorage.storable));

	return {
		subscribe,
		set: (n: T) => {
			isBrowser && (localStorage.setItem(name, JSON.stringify(n)))
			set(n);
		},
		update: (cb: (st: Writable<T>) => T) => {
			const updatedStore = cb(store);

			isBrowser && (localStorage.setItem(name, JSON.stringify(updatedStore)))
			set(updatedStore);
		}
	};
}
