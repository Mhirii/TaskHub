<script lang="ts" module>
	import { z } from "zod/v4";

	const formSchema = z.object({
		username: z.string().min(2).max(50),
		password: z.string().min(8).max(50),
	});
</script>
