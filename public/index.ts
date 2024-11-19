// Follow this setup guide to integrate the Deno language server with your editor:
// https://deno.land/manual/getting_started/setup_your_environment
// This enables autocomplete, go to definition, etc.

// Setup type definitions for built-in Supabase Runtime APIs
import "jsr:@supabase/functions-js/edge-runtime.d.ts";
import { createClient } from "jsr:@supabase/supabase-js@2";
import { serve } from "https://deno.land/std@0.114.0/http/server.ts";

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

const supabase = createClient(
  Deno.env.get("SUPABASE_URL") ?? "",
  Deno.env.get("SUPABASE_ANON_KEY") ?? "",
);

const { data, error } = await supabase.from("vscode_downloads").select("*", {
  count: "exact",
}).order("created_at", { ascending: true });
if (error) {
  console.error(`Failed to save install counts:`, error);
}

console.log(data);
const apexChartScript = `
  <script src="https://cdn.jsdelivr.net/npm/apexcharts"></script>
`;
const apexChartSetup = `
  <script>
    const data = ${JSON.stringify(data)};
    const labels = data.map(item => item.created_at);
    const gitguardianValues = data.map(item => item.gitguardian);
    const cycodeValues = data.map(item => item.cycode);
    const mendValues = data.map(item => item.mend2022);
    const snykValues = data.map(item => item.snyk);
    const sapCredentialDiggerValues = data.map(item => item.sap_credentialdigger);

    const options = {
      chart: {
        type: 'line',
        height: 900,
        zoom: {
          type: 'x',
          enabled: true,
          autoScaleYaxis: true
        },
        toolbar: {
          autoSelected: 'zoom'
        }
      },
      series: [
        {
          name: 'GitGuardian Downloads',
          data: gitguardianValues
        },
        {
          name: 'Cycode Downloads',
          data: cycodeValues
        },
        {
          name: 'Mend Downloads',
          data: mendValues
        },
        {
          name: 'Snyk Downloads',
          data: snykValues
        },
        {
          name: 'SAP Credentialdigger Downloads',
          data: sapCredentialDiggerValues
        }
      ],
      xaxis: {
        categories: labels
      }
    };

    const chart = new ApexCharts(document.querySelector("#myChart"), options);
    chart.render();
  </script>
`;

const handler = async (req: Request): Promise<Response> => {
  const html = `
    <!DOCTYPE html>
    <html lang="en">
    <head>
      <meta charset="UTF-8">
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <title>VSCode Downloads Chart</title>
      ${apexChartScript}
    </head>
    <body>
      <div id="myChart" style="height: 900px; width: 100%;"></div>
      ${apexChartSetup}
    </body>
    </html>
  `;
  return new Response(html, {
    headers: { "Content-Type": "text/html" },
  });
};

serve(handler, { headers: { "Content-Type": "text/html" } });
