# Harmony - A Spotify Client App for Generating Playlists from Billboard Charts

Harmony is a Spotify client application built using Golang. It automatically generates playlists based on current Billboard charts, providing users with up-to-date music recommendations.

## Features

- Fetches current Billboard chart data.
- Creates playlists on Spotify based on Billboard charts.
- Search and adds tracks to the generated playlists.

## Usage

To use Harmony, follow these steps:

1. **Prerequisites**: Ensure you have Go (Golang) installed on your system.

2. **Installation**: Clone the repository and navigate to the project directory:

    ```bash
    git clone https://github.com/yourusername/harmony.git
    cd harmony
    ```

3. **Set up Spotify Developer Credentials**: Follow the instructions in `spotifyclient/README.md` to set up your Spotify Developer credentials.

4. **Build the Project**: Run the following command to build the project:

    ```bash
    go build
    ```

5. **Run Harmony**: Execute Harmony to generate a playlist based on the current Billboard Hot 100 chart:

    ```bash
    ./harmony
    ```

This command initiates the process of fetching Billboard chart data, creating a Spotify playlist titled "Billboard Hot 100", and adding tracks to it.

## Contributing

Contributions are welcome! If you have suggestions, enhancements, or bug fixes, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
