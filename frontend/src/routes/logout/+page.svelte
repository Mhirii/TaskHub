<!-- this is a logout page, please take the look and feel of the login page for -->
<!-- reference, this page should tell the user youve logged out successfully, redirecting to login in {counter of 5 seconds} seconds. -->
<script lang="ts">
	import { onMount } from "svelte";
	import { goto } from "$app/navigation";
	import {
		Card,
		CardContent,
		CardHeader,
		CardTitle,
	} from "$lib/components/ui/card";
	import { Button } from "$lib/components/ui/button";
	import { Unlock } from "@lucide/svelte";

	let counter = 3;

	onMount(() => {
		const interval = setInterval(() => {
			counter--;
			if (counter <= 0) clearInterval(interval);
		}, 1000);

		return () => clearInterval(interval);
	});

	onMount(() => {
		const cookies = document.cookie.split(";");
		for (const cookie of cookies) {
			console.log(cookie);
			const [name] = cookie.split("=").map((c) => c.trim());
			if (name) {
				document.cookie = `${name}=; Path=/; Expires=Thu, 01 Jan 1970 00:00:00 GMT; SameSite=Strict`;
			}
		}
		const interval = setInterval(() => {
			counter--;
			if (counter <= 0) {
				clearInterval(interval);
				goto("/auth/login");
			}
		}, 1000);

		return () => clearInterval(interval);
	});
</script>

<div
	class="container mx-auto p-4 sm:p-6 max-w-7xl min-h-[calc(100vh-10rem)] flex items-center justify-center"
>
	<Card class="w-full max-w-md shadow-lg">
		<CardHeader>
			<CardTitle class="text-lg lg:text-2xl font-semibold">Logged Out</CardTitle
			>
		</CardHeader>
		<CardContent class="text-center space-y-4">
			<div class="text-muted-foreground text-balance">
				You have been logged out successfully.
			</div>
			<div class="text-muted-foreground text-balance">
				Redirecting to login in {counter} second{counter === 1 ? "" : "s"}...
			</div>
			<Button
				variant="outline"
				onclick={() => goto("/auth/login")}
				class="flex items-center gap-2 mx-auto"
			>
				<Unlock />
				Go to Login Now
			</Button>
		</CardContent>
	</Card>
</div>
