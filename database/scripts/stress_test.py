import requests
import json
import time
import concurrent.futures

GATEWAY_URL = "http://localhost:8080/ingest" # Change to GKE LoadBalancer IP if deployed

def send_signal(i):
    data = {
        "type": "STRESS_SIGNAL_V3",
        "payload": {"index": i, "status": "active", "load": "high"}
    }
    try:
        response = requests.post(GATEWAY_URL, json=data, timeout=2)
        print(f"Signal {i}: {response.status_code} - TraceID: {response.json().get('id')}")
    except Exception as e:
        print(f"Signal {i}: Failed - {e}")

# Simulate 50 concurrent agents hitting the nervous system
with concurrent.futures.ThreadPoolExecutor(max_workers=10) as executor:
    executor.map(send_signal, range(50))
