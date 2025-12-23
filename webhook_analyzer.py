import json
from http.server import HTTPServer, BaseHTTPRequestHandler
from datetime import datetime
import sys

# ANSI Farben
class Colors:
    RESET = "\033[0m"
    RED = "\033[31m"
    GREEN = "\033[32m"
    YELLOW = "\033[33m"
    BLUE = "\033[34m"
    PURPLE = "\033[35m"
    CYAN = "\033[36m"
    WHITE = "\033[37m"

class WebhookHandler(BaseHTTPRequestHandler):
    def log_message(self, format, *args):
        pass  # Unterdrücke Standard-Logging

    def handle_request(self):
        timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")

        print(f"\n{Colors.CYAN}{'=' * 60}{Colors.RESET}")
        print(f"{Colors.GREEN}[{timestamp}]{Colors.RESET} Neue Anfrage empfangen")
        print(f"{Colors.CYAN}{'=' * 60}{Colors.RESET}")

        # Method und Path
        print(f"\n{Colors.YELLOW}Methode:{Colors.RESET}  {Colors.GREEN}{self.command}{Colors.RESET}")
        print(f"{Colors.YELLOW}Path:{Colors.RESET}     {Colors.WHITE}{self.path}{Colors.RESET}")
        print(f"{Colors.YELLOW}Von:{Colors.RESET}      {Colors.WHITE}{self.client_address[0]}:{self.client_address[1]}{Colors.RESET}")

        # Headers
        print(f"\n{Colors.PURPLE}--- Headers ---{Colors.RESET}")
        for key, value in sorted(self.headers.items()):
            print(f"{Colors.BLUE}{key}:{Colors.RESET} {value}")

        # Body lesen
        content_length = self.headers.get('Content-Length')
        if content_length:
            body = self.rfile.read(int(content_length))
            if body:
                print(f"\n{Colors.PURPLE}--- Body ({len(body)} Bytes) ---{Colors.RESET}")
                try:
                    json_data = json.loads(body.decode('utf-8'))
                    print(f"{Colors.WHITE}{json.dumps(json_data, indent=2, ensure_ascii=False)}{Colors.RESET}")
                except (json.JSONDecodeError, UnicodeDecodeError):
                    print(f"{Colors.WHITE}{body.decode('utf-8', errors='replace')}{Colors.RESET}")
        else:
            print(f"\n{Colors.PURPLE}--- Kein Body ---{Colors.RESET}")

        print(f"\n{Colors.CYAN}{'-' * 60}{Colors.RESET}")

        # Antwort senden
        self.send_response(200)
        self.send_header('Content-Type', 'application/json')
        self.end_headers()
        response = {
            "status": "received",
            "timestamp": timestamp,
            "message": "Webhook erfolgreich empfangen"
        }
        self.wfile.write(json.dumps(response).encode('utf-8'))

    def do_GET(self):
        self.handle_request()

    def do_POST(self):
        self.handle_request()

    def do_PUT(self):
        self.handle_request()

    def do_DELETE(self):
        self.handle_request()

    def do_PATCH(self):
        self.handle_request()

    def do_OPTIONS(self):
        self.handle_request()

def main():
    # Windows: ANSI Farben aktivieren
    if sys.platform == 'win32':
        import os
        os.system('')

    port = 9999
    server = HTTPServer(('0.0.0.0', port), WebhookHandler)

    print(f"{Colors.CYAN}========================================{Colors.RESET}")
    print(f"{Colors.GREEN}  Webhook Analyzer Server{Colors.RESET}")
    print(f"{Colors.CYAN}========================================{Colors.RESET}")
    print(f"{Colors.WHITE}Server lauscht auf Port {Colors.YELLOW}{port}{Colors.RESET}")
    print(f"{Colors.WHITE}URL: {Colors.YELLOW}http://localhost:{port}{Colors.RESET}")
    print(f"\n{Colors.PURPLE}Warte auf eingehende Webhooks...{Colors.RESET}")
    print(f"{Colors.WHITE}Druecke Ctrl+C zum Beenden{Colors.RESET}")
    print("-" * 50)

    try:
        server.serve_forever()
    except KeyboardInterrupt:
        print(f"\n{Colors.YELLOW}Server beendet.{Colors.RESET}")
        server.shutdown()

if __name__ == "__main__":
    main()
