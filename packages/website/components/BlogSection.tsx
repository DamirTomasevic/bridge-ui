const posts = [
  {
    title: "ZK-Roller-Coaster #8",
    href: "https://taiko.mirror.xyz/tOUCZgLRV9bKH24bxhahISpdhkQmqVyM-ZX-wMWtqkI",
    description:
      "This is the 8th edition of ZK-Roller-Coaster where we track and investigate the most exciting, meaningful, and crazy ZK-stuff of the prior two weeks.",
    date: "Jul 04, 2023",
    datetime: "2023-07-04",
    imageUrl:
      "https://mirror-media.imgix.net/publication-images/3cfm8O9yVJ8aszk8bQ700.png?height=1536&width=3072&h=1536&w=3072&auto=compress",
    readingTime: "4 min",
    author: {
      name: "Lisa A.",
      imageUrl: "https://avatars.githubusercontent.com/u/106527861?v=4",
    },
  },
  {
    title: "ZK-Roller-Coaster #7",
    href: "https://taiko.mirror.xyz/6WL5I2lbpYxOjhU82eUOyUvYa0yF2_rekI0f7cBrGxw",
    description:
      "This is the 7th edition of ZK-Roller-Coaster where we track and investigate the most exciting, meaningful, and crazy ZK-stuff of the prior two weeks.",
    date: "Jun 20, 2023",
    datetime: "2023-06-20",
    imageUrl:
      "https://mirror-media.imgix.net/publication-images/HRLmI4Vmn9A637fxJm8cq.png?height=426&width=851&h=426&w=851&auto=compress",
    readingTime: "4 min",
    author: {
      name: "Lisa A.",
      imageUrl: "https://avatars.githubusercontent.com/u/106527861?v=4",
    },
  },
  {
    title: "L2 MEV wat",
    href: "https://taiko.mirror.xyz/VjNjFws6OOVez5YCDMwjy4BUiDqZBHYDvcW4-JZGDkc",
    description:
      "In this article, we “map” the current landscape of L2 MEV, thinking about different MEV consequences for different L2 designs. We also briefly overview different ways of L2s decentralization and how it might impact L2 MEV.",
    date: "Jun 13, 2023",
    datetime: "2023-06-13",
    imageUrl:
      "https://mirror-media.imgix.net/publication-images/Qgm0gbwbCQnU8bm5Y1dGB.png?height=512&width=1024&h=512&w=1024&auto=compress",
    readingTime: "15 min",
    author: {
      name: "Lisa A.",
      imageUrl: "https://avatars.githubusercontent.com/u/106527861?v=4",
    },
  },
];

export default function BlogSection() {
  return (
    <div className="relative bg-neutral-50 px-4 pt-16 pb-20 sm:px-6 lg:px-8 lg:pt-24 lg:pb-28 dark:bg-neutral-900">
      <div className="absolute inset-0">
        <div className="h-1/3 bg-white sm:h-2/3 dark:bg-neutral-900" />
      </div>
      <div className="relative mx-auto max-w-7xl">
        <div className="text-center">
          <h2 className="font-grotesk text-3xl tracking-tight text-neutral-900 sm:text-4xl dark:text-neutral-100">
            Latest blog posts
          </h2>
          <div className="mx-auto mt-3 max-w-2xl text-xl text-neutral-500 sm:mt-4 dark:text-neutral-300">
            Check out the full blog at{" "}
            <a
              className="underline"
              href="https://taiko.mirror.xyz"
              target="_blank"
              rel="noopener noreferrer"
            >
              taiko.mirror.xyz
            </a>
          </div>
        </div>
        <div className="mx-auto mt-12 grid max-w-lg gap-5 lg:max-w-none lg:grid-cols-3">
          {posts.map((post) => (
            <a
              key={post.title}
              href={post.href}
              target="_blank"
              rel="noopener noreferrer"
              className="hover:shadow-lg transition duration-300"
            >
              <div className="flex flex-col h-full overflow-hidden rounded-lg shadow-lg">
                <div className="flex-shrink-0">
                  <img
                    className="h-54 w-full object-cover"
                    src={post.imageUrl}
                    alt=""
                  />
                </div>
                <div className="flex flex-1 flex-col justify-between bg-white p-6 dark:bg-neutral-800 dark:hover:bg-neutral-700">
                  <div className="flex-1">
                    <div className="mt-2 block">
                      <div className="text-xl font-semibold text-neutral-900 dark:text-neutral-200 line-clamp-1">
                        {post.title}
                      </div>
                      <div className="mt-3 text-base text-neutral-500 dark:text-neutral-300 line-clamp-3">
                        {post.description}
                      </div>
                    </div>
                  </div>
                  <div className="mt-6 flex items-center">
                    <div className="flex-shrink-0">
                      <span className="sr-only">{post.author.name}</span>
                      <img
                        className="h-10 w-10 rounded-full"
                        src={post.author.imageUrl}
                        alt=""
                      />
                    </div>
                    <div className="ml-3">
                      <div className="text-sm font-medium text-[#e81899]">
                        {post.author.name}
                      </div>
                      <div className="flex space-x-1 text-sm text-neutral-500 dark:text-neutral-400">
                        <time dateTime={post.datetime}>{post.date}</time>
                        <span aria-hidden="true">&middot;</span>
                        <span>{post.readingTime} read</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </a>
          ))}
        </div>
      </div>
    </div>
  );
}
