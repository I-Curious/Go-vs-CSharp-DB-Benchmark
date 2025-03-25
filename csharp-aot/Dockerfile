# Use the official .NET SDK image for building the app
FROM mcr.microsoft.com/dotnet/sdk:9.0 AS build
WORKDIR /source

# Install Clang and other required dependencies
RUN apt-get update && \
    apt-get install -y clang gcc make && \
    rm -rf /var/lib/apt/lists/*

# Copy project file and restore dependencies
COPY csharp-aot.csproj ./
RUN dotnet add package Npgsql
RUN dotnet add package Dapper
RUN dotnet restore

# Copy all the source code
COPY . .

# Publish the app using AOT
RUN dotnet publish -c Release -o /app /p:PublishAot=true /p:SelfContained=true /p:StripSymbols=true

# Use a runtime image that supports AOT
FROM mcr.microsoft.com/dotnet/runtime:9.0 AS runtime
WORKDIR /app
COPY --from=build /app .

ENTRYPOINT ["./csharp-aot"]