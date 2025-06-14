<script lang="ts">
	import {
		Card,
		CardHeader,
		CardTitle,
		CardContent,
	} from "$lib/components/ui/card";
	import { Skeleton } from "$lib/components/ui/skeleton";
</script>

<div class="grid gap-4 md:grid-cols-3">
	{#each ["Planned", "In Progress", "Done"] as status}
		<Card>
			<CardHeader>
				<CardTitle class="text-lg">{status}</CardTitle>
			</CardHeader>
			<CardContent>
				<div class="space-y-2">
					<Skeleton class="h-4 w-full" />
					<Skeleton class="h-4 w-full" />
				</div>
			</CardContent>
		</Card>
	{/each}
</div>
