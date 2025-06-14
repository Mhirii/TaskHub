<script lang="ts">
	import { Card, CardContent } from "$lib/components/ui/card";
	import type { Task } from "$lib/types/tasks";

	export let task: Task;
</script>

<Card class="mb-2">
	<CardContent class="p-4">
		<p class="font-medium">{task.name}</p>
		<p class="text-sm text-muted-foreground">{task.description}</p>
	</CardContent>
</Card>
