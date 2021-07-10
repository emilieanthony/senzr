import { UserButton } from '@clerk/clerk-react';
import './header.css';
const Header = () => {
  return (
    <div className="header">
      Header
      <div className="settings">
        <UserButton />
      </div>
    </div>
  )
}

export default Header;
