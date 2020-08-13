# https://hub.docker.com/_/microsoft-dotnet-core
FROM mcr.microsoft.com/dotnet/core/sdk:3.1 AS build
WORKDIR /source

# copy csproj and restore as distinct layers
COPY *.sln .
COPY aspnetapp/*.csproj ./aspnetapp/
RUN dotnet restore -r linux-x64

# copy everything else and build app
COPY aspnetapp/. ./aspnetapp/
WORKDIR /source/aspnetapp
RUN dotnet add package OpenTelemetry.Exporter.Console -v 0.4.0-beta.2
RUN dotnet add package OpenTelemetry.Exporter.OpenTelemetryProtocol -v 0.4.0-beta.2
RUN dotnet restore -r linux-x64
RUN dotnet publish -c release -o /app -r linux-x64 --self-contained false --no-restore

# final stage/image
FROM mcr.microsoft.com/dotnet/core/aspnet:3.1-buster-slim
WORKDIR /app
COPY --from=build /app ./
ENTRYPOINT ["./aspnetapp"]
