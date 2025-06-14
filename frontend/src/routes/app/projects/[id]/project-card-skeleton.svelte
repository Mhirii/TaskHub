<script lang="ts">
	import { Card, CardHeader, CardContent } from "$lib/components/ui/card";
	import { Skeleton } from "$lib/components/ui/skeleton";
</script>

<Card>
	<CardHeader>
		<Skeleton class="h-6 w-1/2" />
		<Skeleton class="h-4 w-3/4" />
	</CardHeader>
	<CardContent>
		<div class="space-y-2">
			<Skeleton class="h-4 w-full" />
			<Skeleton class="h-4 w-full" />
			<Skeleton class="h-4 w-full" />
		</div>
	</CardContent>
</Card>
