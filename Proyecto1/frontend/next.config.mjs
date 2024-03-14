/** @type {import('next').NextConfig} */
const nextConfig = {
  async rewrites() {
    return [
      {
        source: '/api/:path*',
        destination: 'http://192.168.175.128:8080/api/:path*',
      },
    ];
  },
  output: "standalone",
};

export default nextConfig;
