<script lang="ts">
	import * as Card from "$lib/components/ui/card/index.js";
	import { cn } from "$lib/utils.js";
	import { formSchema, type FormSchema } from "./schema";
	import { type SuperValidated, type Infer } from "sveltekit-superforms";
	import LoginForm from "./login-form.svelte";

	let { data }: { data: { form: SuperValidated<Infer<FormSchema>> } } =
		$props();
</script>

<div
	class="bg-muted flex min-h-svh flex-col items-center justify-center p-6 md:p-10"
>
	<div class="w-full max-w-sm md:max-w-3xl">
		<div class={cn("flex flex-col gap-6")}>
			<Card.Root class="overflow-hidden p-0">
				<Card.Content class="grid p-0 md:grid-cols-2">
					<div class="bg-muted relative hidden md:block">
						<img
							src="/asian_worman_folders.jpg"
							alt="placeholder"
							class="absolute inset-0 h-full w-full object-cover dark:brightness-[0.2] dark:grayscale"
						/>
					</div>
					<LoginForm {data} />
				</Card.Content>
			</Card.Root>
			<div
				class="text-muted-foreground *:[a]:hover:text-primary *:[a]:underline *:[a]:underline-offset-4 text-balance text-center text-xs"
			>
				By clicking continue, you agree to our <a href="##">Terms of Service</a>
				and
				<a href="##">Privacy Policy</a>.
			</div>
		</div>
	</div>
</div>
