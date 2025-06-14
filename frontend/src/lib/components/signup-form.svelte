<script lang="ts">
	import * as Card from "$lib/components/ui/card/index.js";
	import { Label } from "$lib/components/ui/label/index.js";
	import { Input } from "$lib/components/ui/input/index.js";
	import { Button } from "$lib/components/ui/button/index.js";
	import { cn } from "$lib/utils.js";
	import type { HTMLAttributes } from "svelte/elements";

	let { class: className, ...restProps }: HTMLAttributes<HTMLDivElement> = $props();

	const id = $props.id();
</script>

<div class={cn("flex flex-col gap-6", className)} {...restProps}>
	<Card.Root class="overflow-hidden p-0">
		<Card.Content class="grid p-0 md:grid-cols-2">
			<div class="bg-muted relative hidden md:block">
				<img
					src="/asian_worman_folders.jpg"
					alt="placeholder"
					class="absolute inset-0 h-full w-full object-cover dark:brightness-[0.2] dark:grayscale"
				/>
			</div>
			<form class="p-6 md:p-8">
				<div class="flex flex-col gap-6">
					<div class="flex flex-col items-center text-center">
						<h1 class="text-2xl font-bold">Welcome back</h1>
						<p class="text-muted-foreground text-balance">
							Login to your TaskHub account
						</p>
					</div>
					<div class="grid gap-3">
						<Label for="email-{id}">Email</Label>
						<Input id="email-{id}" type="email" placeholder="m@example.com" required />
					</div>
					<div class="grid gap-3">
						<Label for="email-{id}">Username</Label>
						<Input id="username-{id}" type="text" placeholder="John Snow" required />
					</div>
					<div class="grid gap-3">
						<div class="flex items-center">
							<Label for="password">Password</Label>
							<a href="##" class="ml-auto text-sm underline-offset-2 hover:underline">
								Forgot your password?
							</a>
						</div>
						<Input id="password-{id}" type="password" required />
					</div>
					<Button type="submit" class="w-full">Login</Button>
					<div class="text-center text-sm">
						Already have an account?
						<a href="/auth/login" class="underline underline-offset-4"> Login in </a>
					</div>
				</div>
			</form>
		</Card.Content>
	</Card.Root>
	<div
		class="text-muted-foreground *:[a]:hover:text-primary *:[a]:underline *:[a]:underline-offset-4 text-balance text-center text-xs"
	>
		By clicking continue, you agree to our <a href="##">Terms of Service</a> and
		<a href="##">Privacy Policy</a>.
	</div>
</div>
