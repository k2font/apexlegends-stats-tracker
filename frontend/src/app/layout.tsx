import './globals.css'

export const metadata = {
  title: 'Apex Legends Tracker',
  description: 'This application collects information about Apex Legends.',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="ja">
      <body>{children}</body>
    </html>
  )
}
