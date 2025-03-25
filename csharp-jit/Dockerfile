# Build stage
FROM mcr.microsoft.com/dotnet/sdk:9.0 AS build
WORKDIR /source

# Copy only .csproj and restore dependencies (cache layer)
COPY csharp-jit.csproj ./
RUN dotnet restore
RUN dotnet add package Npgsql
RUN dotnet add package Dapper

# Copy the rest of the application files
COPY . . 

# Run the build process
RUN dotnet publish -c Release -o /app

# Runtime stage
FROM mcr.microsoft.com/dotnet/aspnet:9.0 AS runtime
WORKDIR /app
COPY --from=build /app .
ENTRYPOINT ["dotnet", "csharp-jit.dll"]