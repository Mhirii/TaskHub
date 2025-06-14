<script lang="ts">
	import { Card, CardContent } from "$lib/components/ui/card";
	import { Button } from "$lib/components/ui/button";
</script>

<Card>
	<CardContent class="text-center py-10">
		<p class="text-muted-foreground">No tasks found. Create your first task!</p>
		<Button class="mt-4">Create Task</Button>
	</CardContent>
</Card>
