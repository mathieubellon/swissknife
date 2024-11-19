vscode:
	deno run --allow-read --env-file --allow-net --allow-env supabase/functions/vscode/index.ts

chart:
	deno run --allow-read --env-file --allow-net --allow-env --watch public/index.ts