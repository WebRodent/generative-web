# Generative Web
This API serves as a backend for the Generative Web project. 
It serves AI generated webpages based on a users input.

It uses a template process to store generic frontend templates with placeholders for the AI generated content.
The API then replaces the placeholders with the generated content and serves the page to the user.

## Installation
1. Clone the repository
2. Open devcontainer
3. Run `air` to start the server

## Usage
1. Send a POST request to `/generate` with the following body:
```json
{
    "template": "template_name",
    "content": {
        "tags": ["tag1", "tag2", "tag3"],
    }
}
```