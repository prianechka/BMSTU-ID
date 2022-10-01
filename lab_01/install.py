from sys import platform
from subprocess import check_output
import hashlib


def get_system_hash():
    result = ""
    if platform == "linux":
        system_uuid = check_output("dmidecode -s system-uuid", shell=True).decode()
        serial_num = check_output("dmidecode -s system-serial-number", shell=True).decode()
        check_str = system_uuid + "_" + serial_num
        result = hashlib.sha256(check_str.encode('utf-8')).hexdigest()
    return result

def check_hash():
    real_key = read_license_key()
    return real_key == get_system_hash()

def write_license_key(got_hash):
    with open("license.key", "w") as licence_file:
        licence_file.write(str(got_hash))

def read_license_key():
    with open("license.key", "r") as lic_file:
        return lic_file.readline()

if __name__ == "__main__":
    write_license_key(get_system_hash())