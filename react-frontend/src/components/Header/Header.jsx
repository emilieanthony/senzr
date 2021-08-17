import { UserButton } from '@clerk/clerk-react';
import './header.css';
const Header = () => {
  return (
    <div className="header">
      <h2 className="title">Senzr</h2>
      <div className="settings">
        <UserButton />
      </div>
    </div>
  )
}

export default Header;
