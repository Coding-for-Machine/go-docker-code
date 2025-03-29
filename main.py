import requests

url = "http://ip172-18-0-27-cumf5nq91nsg009pts3g-3000.direct.labs.play-with-docker.com/run-test"
headers = {"Content-Type": "application/json"}
data = {
    "user_code": "print(input())",
    "language": "python",
    "test_cases": [  # Diqqat, test_case emas, test_cases!
        {
            "test_case": 1,
            "input": "Hello",
            "expected_output": "Hello",
            "input_type": "string",
            "expected_output_type": "string"
        },
        {
            "test_case": 2,
            "input": "123",
            "expected_output": "123",
            "input_type": "integer",
            "expected_output_type": "integer"
        }
    ]
}

response = requests.post(url, json=data, headers=headers)
print(response.json())

