import { Button } from "@/components/ui/button";
import DropdownAvatar from "../Header/DropdownAvatar";

export const Header = () => {
  // TODO: Implement user authentication
  const loggedIn = true;

  return (
    <div className="flex justify-between px-8 w-screen h-16 items-center drop-shadow-2xl shadow-sm">
      <h1 className="font-syuku font-bold text-3xl">WebTech 道場</h1>
      {loggedIn ? (
        // TODO: Implement user authentication
        <DropdownAvatar imgSrc="https://github.com/shadcn.png" fallback="name" />
      ) : (
        <Button>ログイン</Button>
      )
      }
    </div>
  );
};