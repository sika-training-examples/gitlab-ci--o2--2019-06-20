import unittest
import requests

class Test(unittest.TestCase):
    def test_response_code(self):
        response = requests.get('http://server')
        self.assertEqual(200, response.status_code)