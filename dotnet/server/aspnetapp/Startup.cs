using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.HttpsPolicy;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using OpenTelemetry;
using OpenTelemetry.Trace;
using OpenTelemetry.Resources;
using Grpc.Core;

namespace aspnetapp
{
    public class Startup
    {
        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        public IConfiguration Configuration { get; }

        // This method gets called by the runtime. Use this method to add services to the container.
        public void ConfigureServices(IServiceCollection services)
        {
            services.AddRazorPages();
            services.AddOpenTelemetryTracing((builder) => builder
                .AddAspNetCoreInstrumentation(opt => {
                    opt.Propagator = new OpenTelemetry.Context.Propagation.B3Propagator();
                })
                .AddOtlpExporter(opt => {
                    opt.Endpoint = Environment.GetEnvironmentVariable("OTEL_EXPORTER_OTLP_SPAN_ENDPOINT");
                    opt.Headers = new Metadata
                    {
                        { "lightstep-access-token", Environment.GetEnvironmentVariable("LS_ACCESS_TOKEN")}
                    };
                    opt.Credentials = new SslCredentials();
                })
                .SetResource(Resources.CreateServiceResource(Environment.GetEnvironmentVariable("LS_SERVICE_NAME"), serviceVersion: Environment.GetEnvironmentVariable("LS_SERVICE_VERSION")))
            );
        }

        // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
            }
            else
            {
                app.UseExceptionHandler("/Error");
                // The default HSTS value is 30 days. You may want to change this for production scenarios, see https://aka.ms/aspnetcore-hsts.
                app.UseHsts();
            }

            app.UseHttpsRedirection();
            app.UseStaticFiles();

            app.UseRouting();

            app.UseAuthorization();

            app.UseEndpoints(endpoints =>
            {
                endpoints.MapRazorPages();
            });
        }
    }
}
