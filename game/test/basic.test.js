describe('Basic', () => {
  beforeAll(async () => {
    await page.goto('http://127.0.0.1:8080');
  });

  it('title should be "Cloud Journey"', async () => {
    const title = await page.title();
    expect(title).toContain('Cloud Journey');
  });

  it('privacy link must be present"', async () => {
    await expect(page).toMatch('Privacy');
  });

  it('terms link must be present"', async () => {
    await expect(page).toMatch('Terms');
  });

});
