// Follow this setup guide to integrate the Deno language server with your editor:
// https://deno.land/manual/getting_started/setup_your_environment
// This enables autocomplete, go to definition, etc.

// Setup type definitions for built-in Supabase Runtime APIs
import "jsr:@supabase/functions-js/edge-runtime.d.ts";
import { createClient } from "jsr:@supabase/supabase-js@2";
import { DOMParser } from "https://deno.land/x/deno_dom@v0.1.48/deno-dom-wasm.ts";

/* To invoke locally:

  1. Run `supabase start` (see: https://supabase.com/docs/reference/cli/supabase-start)
  2. Make an HTTP request:

  curl -i --location --request POST 'http://127.0.0.1:54321/functions/v1/vscode' \
    --header 'Authorization: Bearer ' \
    --header 'Content-Type: application/json' \
    --data '{"name":"Functions"}'

*/

interface Competitor {
  name: string;
  url: string;
}

const competitors: Competitor[] = [
  {
    name: "gitguardian",
    url:
      "https://marketplace.visualstudio.com/items?itemName=gitguardian-secret-security.gitguardian",
  },
  {
    name: "cycode",
    url: "https://marketplace.visualstudio.com/items?itemName=cycode.cycode",
  },
  {
    name: "snyk",
    url:
      "https://marketplace.visualstudio.com/items?itemName=snyk-security.snyk-vulnerability-scanner-vs",
  },
  {
    name: "mend2019",
    url: "https://marketplace.visualstudio.com/items?itemName=Mend.mend-vs2019",
  },
  {
    name: "mend2022",
    url: "https://marketplace.visualstudio.com/items?itemName=Mend.mend-vs2022",
  },
  {
    name: "sap_credentialdigger",
    url:
      "https://marketplace.visualstudio.com/items?itemName=SAPOSS.vs-code-extension-for-project-credential-digger",
  },
];

const supabase = createClient(
  Deno.env.get("SUPABASE_URL") ?? "",
  Deno.env.get("SUPABASE_ANON_KEY") ?? "",
);

interface CompetitorCount {
  name: string;
  count: number;
}

interface payload {
  [key: string]: number;
}

const payload: payload = {};

await Promise.all(
  competitors.map(async (competitor) => {
    try {
      const count = await getInstallCount(competitor.url);
      payload[competitor.name] = count;
    } catch (error) {
      console.error(
        `Failed to get install count for ${competitor.name}:`,
        error,
      );
    }
  }),
);

const { error } = await supabase.from("vscode_downloads").insert(payload);
if (error) {
  console.error(`Failed to save install counts:`, error);
}

async function getInstallCount(url: string): Promise<number> {
  const response = await fetch(url);
  if (!response.ok) {
    throw new Error(
      `Failed to retrieve the page. Status code: ${response.status}`,
    );
  }

  const text = await response.text();
  const parser = new DOMParser();
  const doc = parser.parseFromString(text, "text/html");

  const installsText = doc.querySelector("span.installs-text")?.textContent;
  if (!installsText) {
    throw new Error("The specified element was not found");
  }

  const countStr = installsText.split(" ")[1].replace(/,/g, "");
  const count = parseInt(countStr, 10);
  if (isNaN(count)) {
    throw new Error("Failed to convert install count to integer");
  }

  return count;
}
