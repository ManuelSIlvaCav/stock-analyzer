export default function CompanyLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <>
      <div>layout</div>
      <main>{children}</main>
    </>
  );
}
