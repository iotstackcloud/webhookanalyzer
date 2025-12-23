# Webhook Analyzer

Ein einfacher Webserver zur Analyse von eingehenden Webhooks. Zeigt alle Details der HTTP-Anfragen im Terminal an.

## Features

- Lauscht auf Port **9999**
- Unterstützt alle HTTP-Methoden (GET, POST, PUT, DELETE, PATCH, OPTIONS)
- Farbige Terminal-Ausgabe
- Automatische JSON-Formatierung
- Zeigt Headers, Body und Metadaten an
- Sendet JSON-Bestätigung zurück

## Installation

### Python

```bash
python webhook_analyzer.py
```

### Als EXE (Windows)

```bash
pip install pyinstaller
python -m PyInstaller --onefile --console --name webhook_analyzer webhook_analyzer.py
```

Die EXE befindet sich dann in `dist/webhook_analyzer.exe`.

## Verwendung

1. Server starten:
   ```bash
   webhook_analyzer.exe
   # oder
   python webhook_analyzer.py
   ```

2. Webhooks an `http://localhost:9999` senden:
   ```bash
   curl -X POST http://localhost:9999/test \
     -H "Content-Type: application/json" \
     -d '{"event": "test", "data": {"key": "value"}}'
   ```

## Beispiel-Ausgabe

```
========================================
  Webhook Analyzer Server
========================================
Server lauscht auf Port 9999
URL: http://localhost:9999

Warte auf eingehende Webhooks...
--------------------------------------------------

============================================================
[2024-01-15 14:30:22] Neue Anfrage empfangen
============================================================

Methode:  POST
Path:     /test
Von:      127.0.0.1:54321

--- Headers ---
Content-Length: 42
Content-Type: application/json
Host: localhost:9999

--- Body (42 Bytes) ---
{
  "event": "test",
  "data": {
    "key": "value"
  }
}

------------------------------------------------------------
```

## Lizenz

MIT
