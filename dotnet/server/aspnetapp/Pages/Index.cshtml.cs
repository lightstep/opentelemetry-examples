using System;
using System.Collections.Generic;
using System.Linq;
using System.Diagnostics;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Microsoft.Extensions.Logging;
using OpenTelemetry;
using OpenTelemetry.Resources;
using OpenTelemetry.Trace;
using Grpc.Core;

namespace aspnetapp.Pages
{
    public class IndexModel : PageModel
    {
        static readonly ActivitySource activitySource = new ActivitySource(
        "MyCompany.MyProduct.MyLibrary");
        private readonly ILogger<IndexModel> _logger;

        public IndexModel(ILogger<IndexModel> logger)
        {
            _logger = logger;
        }

        public void OnGet()
        {

            using var otel = Sdk.CreateTracerProvider(b => b
                .AddActivitySource("MyCompany.MyProduct.MyLibrary")
                .UseOtlpExporter(opt =>
                {
                    opt.Endpoint = Environment.GetEnvironmentVariable("OTEL_EXPORTER_OTLP_SPAN_ENDPOINT");
                    opt.Headers = new Metadata
                    {
                        { "lightstep-access-token", Environment.GetEnvironmentVariable("LS_ACCESS_TOKEN")}
                    };
                    opt.Credentials = new SslCredentials();
                })
                .SetResource(Resources.CreateServiceResource(Environment.GetEnvironmentVariable("LS_SERVICE_NAME"), serviceVersion: Environment.GetEnvironmentVariable("LS_SERVICE_VERSION"))));

            using (var activity = activitySource.StartActivity("SayHello"))
            {
                activity?.AddTag("bar", "Hello, World!");
            }

        }
    }
}
