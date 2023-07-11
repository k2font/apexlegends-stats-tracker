import Link from 'next/Link';

const Sidebar = () => {
  return (
    <>
      <div className="sidebar">
        <ul className="sidebarContents">
          <li><Link href="/server-status">Server Status</Link></li>
          <li><Link href="/map-rotation">Map Rotation</Link></li>
          <li><Link href="/predator">Predator Border</Link></li>
          <li><Link href="/search">Search</Link></li>
        </ul>
      </div>
    </>
  )
}

export default Sidebar;