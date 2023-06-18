/** @type {import('next').NextConfig} */
const nextConfig = {
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'apexlegendsstatus.com',
        port: '',
        pathname: '**',
      },
    ],
  },
}

module.exports = nextConfig
